package home

import (
	"context"

	"Gomall/app/frontend/biz/service"
	"Gomall/app/frontend/biz/utils"
	common "Gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 加载特定的模板文件
	c.HTML(consts.StatusOK, "home", utils.WarpResponse(ctx, c, resp))
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
