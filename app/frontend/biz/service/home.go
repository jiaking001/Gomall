package service

import (
	"Gomall/app/frontend/hertz_gen/frontend/common"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "兄弟，你好香", "Price": "100$", "Picture": "/static/image/001.jpg"},
		{"Name": "兄弟，你好香", "Price": "100$", "Picture": "/static/image/001.jpg"},
		{"Name": "兄弟，你好香", "Price": "100$", "Picture": "/static/image/001.jpg"},
		{"Name": "兄弟，你好香", "Price": "100$", "Picture": "/static/image/001.jpg"},
		{"Name": "兄弟，你好香", "Price": "100$", "Picture": "/static/image/001.jpg"},
		{"Name": "兄弟，你好香", "Price": "100$", "Picture": "/static/image/001.jpg"},
		{"Name": "兄弟，你好香", "Price": "100$", "Picture": "/static/image/001.jpg"},
		{"Name": "兄弟，你好香", "Price": "100$", "Picture": "/static/image/001.jpg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
