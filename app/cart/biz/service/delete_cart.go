package service

import (
	"Gomall/app/cart/biz/dal/mysql"
	"Gomall/app/cart/biz/model"
	cart "Gomall/rpc_gen/kitex_gen/cart"
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type DeleteCartService struct {
	ctx context.Context
} // NewDeleteCartService new DeleteCartService
func NewDeleteCartService(ctx context.Context) *DeleteCartService {
	return &DeleteCartService{ctx: ctx}
}

// Run create note info
func (s *DeleteCartService) Run(req *cart.DeleteCartReq) (resp *cart.DeleteCartResp, err error) {
	// Finish your business logic.
	err = model.DeleteCart(s.ctx, mysql.DB, req.ProductId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	return &cart.DeleteCartResp{}, nil
}
