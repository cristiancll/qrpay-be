package middleware

import (
	"context"
	"encoding/json"
	"github.com/cristiancll/qrpay-be/configs"
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/cristiancll/qrpay-be/internal/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
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
		return nil, status.Error(codes.Unauthenticated, errors.AUTH_ERROR)
	}

	cookies := md.Get("cookie")
	if len(cookies) == 0 {
		return nil, status.Error(codes.Unauthenticated, errors.AUTH_ERROR)
	}
	cookie := http.Header{"Cookie": []string{cookies[0]}}
	requestCookie, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.AUTH_ERROR)
	}
	requestCookie.Header = cookie
	jwtToken, err := requestCookie.Cookie("jwtToken")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.AUTH_ERROR)
	}
	return jwtToken, nil
}

func getTokenFromAuthorization(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, errors.AUTH_ERROR)
	}

	tokenStrings := md.Get("authorization")
	if len(tokenStrings) == 0 {
		return "", status.Error(codes.Unauthenticated, errors.AUTH_ERROR)
	}
	tokenString := tokenStrings[0][len("Bearer "):]
	return tokenString, nil
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if authlessEndpoint(info) {
		return handler(ctx, req)
	}

	jwtToken, err := getTokenFromCookie(ctx)
	if err != nil {
		return nil, err
	}

	tokenString := jwtToken.Value
	publicKey := configs.Get().Keys.JWT.PublicKey
	claims, refreshedToken, err := security.VerifyJWTToken(tokenString, publicKey)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, errors.AUTH_ERROR)
	}
	if refreshedToken != "" {
		err = security.UpdateJWTCookie(ctx, refreshedToken)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
		}
	}

	var subjClaims security.SubjectClaims
	err = json.Unmarshal([]byte(claims.Subject), &subjClaims)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	ctx = context.WithValue(ctx, "UUID", subjClaims.UUID)
	ctx = context.WithValue(ctx, "Role", subjClaims.Role)
	return handler(ctx, req)
}
