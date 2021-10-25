package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/gofromzero/kratos-temp/api/user/v1"
	"github.com/gofromzero/kratos-temp/app/user/service/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUseCase

	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/server-service"))}
}
