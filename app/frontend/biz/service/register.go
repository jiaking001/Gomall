package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	auth "Gomall/app/frontend/hertz_gen/frontend/auth"
	common "Gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.Register) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// TODO USER SVC API
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	_ = session.Save()
	return
}
