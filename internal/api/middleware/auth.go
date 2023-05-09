package middleware

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/configs"
	"github.com/cristiancll/qrpay-be/internal/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
)

var authByPassList = []string{
	"/" + proto.UserService_ServiceDesc.ServiceName + "/Create",
}

func authBypass(info *grpc.UnaryServerInfo) bool {
	switch info.Server.(type) { // TODO: test this
	case proto.AuthServiceServer:
		return true
	}
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
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	cookies := md.Get("cookie")
	if len(cookies) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "cookie is not provided")
	}
	cookie := http.Header{"Cookie": []string{cookies[0]}}
	requestCookie, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to parse cookie")
	}
	requestCookie.Header = cookie
	jwtToken, err := requestCookie.Cookie("jwtToken")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid cookie")
	}
	return jwtToken, nil
}

func getTokenFromAuthorization(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	tokenStrings := md.Get("authorization")
	if len(tokenStrings) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	tokenString := tokenStrings[0][len("Bearer "):]
	return tokenString, nil
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if authBypass(info) {
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
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %v", err)
	}
	if refreshedToken != "" {
		err = security.UpdateJWTCookie(ctx, refreshedToken)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update cookie: %v", err)
		}
	}

	ctx = context.WithValue(ctx, "UUID", claims.Subject)
	// TODO: add role to claims
	//ctx = context.WithValue(ctx, "Role", claims.Role)
	return handler(ctx, req)
}
