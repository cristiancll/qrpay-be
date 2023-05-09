package security

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/cristiancll/qrpay-be/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GenerateJWTToken(uuid string, privateKey *ecdsa.PrivateKey) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(configs.Get().JWT.GetExpiresIn())),
		Issuer:    configs.Get().JWT.Issuer,
		Subject:   uuid,
		//Role:      role, // TODO: Add role to claims
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
			return nil, errors.New("invalid signing method")
		}

		kid, ok := token.Header["kid"].(string)
		if !ok || kid != configs.Get().JWT.KeyID {
			return nil, errors.New("invalid key id")
		}

		return publicKey, nil
	})

	var refreshedToken string

	if err != nil {
		return nil, refreshedToken, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, refreshedToken, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	now := time.Now().Unix()
	expiresAt := claims.ExpiresAt.Time.Unix()
	refreshTime := expiresAt - int64(configs.Get().JWT.GetExpiresIn().Seconds())

	if now > expiresAt {
		return nil, refreshedToken, status.Error(codes.Unauthenticated, "token has expired")
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
		return status.Errorf(codes.Internal, "Error sending cookie: %v", err)
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
		return status.Errorf(codes.Internal, "Error sending cookie: %v", err)
	}
	return nil
}
