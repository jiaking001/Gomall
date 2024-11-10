package service

import (
	"Gomall/app/frontend/hertz_gen/frontend/common"
	"Gomall/app/frontend/infra/rpc"
	"Gomall/rpc_gen/kitex_gen/product"
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil
}
