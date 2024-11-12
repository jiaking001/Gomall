package user

import (
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUserByUserId(ctx context.Context, req *user.GetUserByUserIdReq, callOptions ...callopt.Option) (resp *user.GetUserByUserIdResp, err error) {
	resp, err = defaultClient.GetUserByUserId(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUserByUserId call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
