package service

import (
	user "Gomall/rpc_gen/kitex_gen/user"
	"context"
	"testing"
)

func TestGetUserById_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetUserByIdService(ctx)
	// init req and assert value

	req := &user.GetUserByUserIdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
