package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/configs"
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

var authByPassList = []string{
	"/" + proto.UserService_ServiceDesc.ServiceName + "/Create",
	"/" + proto.AuthService_ServiceDesc.ServiceName + "/Login",
}

func authlessEndpoint(info *grpc.UnaryServerInfo) bool {
	for _, v := range authByPassList {
		if info.FullMethod == v {
			return true
		}
	}
	return false
}

func getTokenFromCookie(ctx context.Context) (*http.Cookie, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errs.New(errors.New(""), errCode.Internal)
	}

	cookies := md.Get("cookie")
	if len(cookies) == 0 {
		return nil, errs.New(errors.New(""), errCode.Internal)
	}
	cookie := http.Header{"Cookie": []string{cookies[0]}}
	requestCookie, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	requestCookie.Header = cookie
	jwtToken, err := requestCookie.Cookie("jwtToken")
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return jwtToken, nil
}

func extractToken(tokenStrings []string) (string, error) {
	if len(tokenStrings) == 0 {
		return "", fmt.Errorf("no authorization token provided")
	}

	tokenString := strings.TrimSpace(tokenStrings[0])

	var sanitizedToken string
	// Check if the token starts with "Bearer "
	if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
		// Remove the "Bearer " prefix
		sanitizedToken = tokenString[len("bearer "):len(tokenString)]
	}

	// Remove any leading or trailing spaces from the token
	sanitizedToken = strings.TrimSpace(sanitizedToken)

	if sanitizedToken == "" {
		return "", fmt.Errorf("invalid authorization token")
	}

	return sanitizedToken, nil
}

func getTokenFromAuthorization(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errs.New(errors.New(""), errCode.Internal)
	}

	tokenStrings := md.Get("Authorization")
	if len(tokenStrings) == 0 {
		return "", errs.New(errors.New(""), errCode.Internal)
	}
	token, err := extractToken(tokenStrings)
	if err != nil {
		return "", errs.New(err, errCode.Internal)
	}
	return token, nil
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if authlessEndpoint(info) {
		return handler(ctx, req)
	}

	var tokenString string
	var err error
	if configs.Get().JWT.IsSourceCookies() {
		jwtToken, err := getTokenFromCookie(ctx)
		if err != nil {
			return nil, errs.New(err, errCode.Internal)
		}
		tokenString = jwtToken.Value
	} else {
		tokenString, err = getTokenFromAuthorization(ctx)
		if err != nil {
			return nil, errs.New(err, errCode.Internal)
		}
	}

	publicKey := configs.Get().Keys.JWT.PublicKey
	claims, refreshedToken, err := security.VerifyJWTToken(tokenString, publicKey)
	if err != nil {
		return nil, errs.Wrap(errors.New(""), "")
	}
	if refreshedToken != "" {
		if configs.Get().JWT.IsSourceCookies() {
			err = security.UpdateJWTCookie(ctx, refreshedToken)
			if err != nil {
				return nil, errs.New(err, errCode.Internal)
			}
		} else {
			ctx = context.WithValue(ctx, "RefreshedToken", refreshedToken)
		}
	}

	var subjClaims security.SubjectClaims
	err = json.Unmarshal([]byte(claims.Subject), &subjClaims)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	ctx = context.WithValue(ctx, "UUID", subjClaims.UUID)
	ctx = context.WithValue(ctx, "Role", subjClaims.Role)
	return handler(ctx, req)
}
