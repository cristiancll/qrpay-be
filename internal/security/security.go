package security

import (
	"context"
	"crypto/ecdsa"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/configs"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type SubjectClaims struct {
	UUID string `json:"uuid"`
	Role string `json:"role"`
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errs.New(err, errCode.Internal)
	}
	return string(hash), nil
}

func CheckPassword(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errs.New(err, errCode.AccessDenied)
	}
	return nil
}

func GenerateJWTToken(subject string, privateKey *ecdsa.PrivateKey) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(configs.Get().JWT.GetExpiresIn())),
		Issuer:    configs.Get().JWT.Issuer,
		Subject:   subject,
	}

	signingMethod := jwt.GetSigningMethod(configs.Get().JWT.SigningAlgorithm)
	token := jwt.NewWithClaims(signingMethod, claims)
	token.Header["kid"] = configs.Get().JWT.KeyID
	return token.SignedString(privateKey)
}

func VerifyJWTToken(tokenString string, publicKey *ecdsa.PublicKey) (*jwt.RegisteredClaims, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		signingMethod := jwt.GetSigningMethod(configs.Get().JWT.SigningAlgorithm)
		if token.Method != signingMethod {
			return nil, errs.New(errors.New(errMsg.SigningMethodMismatch), errCode.Unauthenticated)
		}

		kid, ok := token.Header["kid"].(string)
		if !ok || kid != configs.Get().JWT.KeyID {
			return nil, errs.New(errors.New(errMsg.KIDMismatch), errCode.Unauthenticated)
		}

		return publicKey, nil
	})

	var refreshedToken string

	if err != nil {
		return nil, refreshedToken, errs.New(err, errCode.Unauthenticated)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, refreshedToken, errs.New(errors.New(errMsg.InvalidToken), errCode.Unauthenticated)
	}

	now := time.Now().Unix()
	expiresAt := claims.ExpiresAt.Time.Unix()
	refreshTime := expiresAt - int64(configs.Get().JWT.GetRefreshThreshold().Seconds())
	if now > expiresAt {
		return nil, refreshedToken, errs.New(errors.New(errMsg.TokenExpired), errCode.Unauthenticated)
	}

	if now > refreshTime {
		newToken, _ := GenerateJWTToken(claims.Subject, configs.Get().Keys.JWT.PrivateKey)
		if newToken != "" {
			refreshedToken = newToken
		}
	}

	return claims, refreshedToken, nil
}

func UpdateJWTCookie(ctx context.Context, newToken string) error {
	cookie := &http.Cookie{
		Name:     "jwtToken",
		Value:    newToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(configs.Get().JWT.GetExpiresIn()),
	}
	headers := metadata.Pairs("Set-Cookie", cookie.String())

	err := grpc.SendHeader(ctx, headers)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	return nil
}

func DeleteJWTCookie(ctx context.Context) error {
	cookie := &http.Cookie{
		Name:     "jwtToken",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}
	headers := metadata.Pairs("Set-Cookie", cookie.String())

	err := grpc.SendHeader(ctx, headers)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	return nil
}

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes  = "0123456789"
	specialBytes = "!@#$%^&*()-_=+[]{}<>?/|"
)

func RandomPassword() string {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	var passwordBuilder strings.Builder

	for i := 0; i < 20; i++ {
		switch i % 4 {
		case 0:
			passwordBuilder.WriteByte(letterBytes[random.Intn(len(letterBytes))])
		case 1:
			passwordBuilder.WriteByte(numberBytes[random.Intn(len(numberBytes))])
		case 2:
			passwordBuilder.WriteByte(specialBytes[random.Intn(len(specialBytes))])
		case 3:
			passwordBuilder.WriteByte(letterBytes[random.Intn(len(letterBytes))])
		}
	}

	return passwordBuilder.String()
}
