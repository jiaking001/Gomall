package service

import (
	"context"

	home "Gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/001.jpg"},
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/001.jpg"},
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/001.jpg"},
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/001.jpg"},
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/001.jpg"},
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/001.jpg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
