package service

import (
	cart "Gomall/app/frontend/hertz_gen/frontend/cart"
	common "Gomall/app/frontend/hertz_gen/frontend/common"
	"Gomall/app/frontend/infra/rpc"
	frontendUtils "Gomall/app/frontend/utils"
	rpccart "Gomall/rpc_gen/kitex_gen/cart"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	userId := uint32(frontendUtils.GetUserIdFromCtx(h.Context))
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: userId,
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  int32(req.ProductNum),
		},
	})
	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	for _, item := range cartResp.Items {
		if item.Quantity == 0 {
			_, e := rpc.CartClient.DeleteCart(h.Context, &rpccart.DeleteCartReq{
				ProductId: item.ProductId,
			})
			if e != nil {
				return nil, e
			}
		}
	}
	return
}
