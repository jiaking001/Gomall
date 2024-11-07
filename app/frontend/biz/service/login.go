package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	auth "Gomall/app/frontend/hertz_gen/frontend/auth"
	common "Gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo USER SVC API

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	_ = session.Save()
	return
}
