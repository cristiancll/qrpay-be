package handler

import (
	"context"
	"encoding/json"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/configs"
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/cristiancll/qrpay-be/internal/security"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type Auth interface {
	Login(ctx context.Context, req *proto.AuthLoginRequest) (*proto.AuthLoginResponse, error)
	Logout(ctx context.Context, req *proto.AuthVoid) (*proto.AuthVoid, error)
	Heartbeat(ctx context.Context, req *proto.AuthVoid) (*proto.AuthHeartbeatResponse, error)
	proto.AuthServiceServer
}

type auth struct {
	service service.Auth
	proto.UnimplementedAuthServiceServer
}

func NewAuth(service service.Auth) Auth {
	return &auth{
		service: service,
	}
}

func (h *auth) Login(ctx context.Context, req *proto.AuthLoginRequest) (*proto.AuthLoginResponse, error) {
	if req.Phone == "" {
		return nil, errs.New(errors.New(errMsg.PhoneRequired), errCode.InvalidArgument)
	}
	if req.Password == "" {
		return nil, errs.New(errors.New(errMsg.PasswordRequired), errCode.InvalidArgument)
	}
	user, auth, err := h.service.Login(ctx, req.Phone, req.Password)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedLogin, req.Phone) // Don't log the password
	}

	privateKey := configs.Get().Keys.JWT.PrivateKey
	subj, err := json.Marshal(security.SubjectClaims{
		UUID: user.UUID,
		Role: strconv.FormatInt(int64(user.Role), 10),
	})
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}

	res := &proto.AuthLoginResponse{
		User: &proto.User{
			Uuid:      user.UUID,
			Name:      user.Name,
			Role:      int64(user.Role),
			Phone:     user.Phone,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
		Auth: &proto.Auth{
			Verified: auth.Verified,
			Disabled: auth.Disabled,
		},
	}

	token, err := security.GenerateJWTToken(string(subj[:]), privateKey)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}

	if configs.Get().JWT.IsSourceCookies() {
		err = security.UpdateJWTCookie(ctx, token)
		if err != nil {
			return nil, errs.New(err, errCode.Internal)
		}
	} else {
		res.Token = &token
	}
	return res, nil
}

func (h *auth) Logout(ctx context.Context, req *proto.AuthVoid) (*proto.AuthVoid, error) {
	err := security.DeleteJWTCookie(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedDeleteCookie)
	}
	res := &proto.AuthVoid{}
	return res, nil
}

func (h *auth) Heartbeat(ctx context.Context, req *proto.AuthVoid) (*proto.AuthHeartbeatResponse, error) {
	user, auth, err := h.service.Heartbeat(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedHeartbeat)
	}
	res := &proto.AuthHeartbeatResponse{
		User: &proto.User{
			Uuid:      user.UUID,
			Name:      user.Name,
			Role:      int64(user.Role),
			Phone:     user.Phone,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
		Auth: &proto.Auth{
			Verified: auth.Verified,
			Disabled: auth.Disabled,
		},
	}

	if !configs.Get().JWT.IsSourceCookies() {
		refreshedToken := ctx.Value("RefreshedToken")
		if refreshedToken != nil {
			token, ok := refreshedToken.(string)
			if !ok {
				return nil, errs.New(errors.New(errMsg.TokenNotString), errCode.Internal, refreshedToken)
			}
			if token != "" {
				res.Token = &token
			}
		}

	}
	return res, nil
}
