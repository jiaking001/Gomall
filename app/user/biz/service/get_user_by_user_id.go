package service

import (
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
)

type GetUserByUserIdService struct {
	ctx context.Context
} // NewGetUserByUserIdService new GetUserByUserIdService
func NewGetUserByUserIdService(ctx context.Context) *GetUserByUserIdService {
	return &GetUserByUserIdService{ctx: ctx}
}

// Run create note info
func (s *GetUserByUserIdService) Run(req *user.GetUserByUserIdReq) (resp *user.GetUserByUserIdResp, err error) {
	// Finish your business logic.

	return
}
