package service

import (
	"Gomall/app/user/biz/dal/mysql"
	"Gomall/app/user/biz/model"
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
)

type GetUserByIdService struct {
	ctx context.Context
} // NewGetUserByIdService new GetUserByIdService
func NewGetUserByIdService(ctx context.Context) *GetUserByIdService {
	return &GetUserByIdService{ctx: ctx}
}

// Run create note info
func (s *GetUserByIdService) Run(req *user.GetUserByUserIdReq) (resp *user.GetUserByUserIdResp, err error) {
	// Finish your business logic.
	row, _ := model.GetUserByUserId(mysql.DB, uint(req.UserId))
	resp = &user.GetUserByUserIdResp{
		User: &user.User{
			Email: row.Email,
		},
	}
	return resp, nil
}
