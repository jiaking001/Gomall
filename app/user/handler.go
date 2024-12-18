package main

import (
	"Gomall/app/user/biz/service"
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// GetUserByUserId implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByUserId(ctx context.Context, req *user.GetUserByUserIdReq) (resp *user.GetUserByUserIdResp, err error) {
	resp, err = service.NewGetUserByIdService(ctx).Run(req)

	return resp, err
}
