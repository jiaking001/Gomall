package utils

import (
	"Gomall/app/frontend/infra/rpc"
	frontendUtils "Gomall/app/frontend/utils"
	"Gomall/rpc_gen/kitex_gen/cart"
	"Gomall/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	content["user_id"] = frontendUtils.GetUserIdFromCtx(ctx)

	if frontendUtils.GetUserIdFromCtx(ctx) > 0 {
		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: uint32(frontendUtils.GetUserIdFromCtx(ctx)),
		})
		if err == nil && cartResp != nil {
			content["cart_num"] = len(cartResp.Items)
		}
	}

	if frontendUtils.GetUserIdFromCtx(ctx) > 0 {
		userResp, err := rpc.UserClient.GetUserByUserId(ctx, &user.GetUserByUserIdReq{
			UserId: uint32(frontendUtils.GetUserIdFromCtx(ctx)),
		})
		if err == nil && userResp != nil {
			content["email"] = userResp.User.Email
		}
	}

	return content
}
