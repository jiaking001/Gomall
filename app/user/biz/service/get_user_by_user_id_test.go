package service

import (
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
	"testing"
)

func TestGetUserByUserId_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetUserByUserIdService(ctx)
	// init req and assert value

	req := &user.GetUserByUserIdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
