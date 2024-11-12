package service

import (
	cart "Gomall/rpc_gen/kitex_gen/cart"
	"context"
	"testing"
)

func TestDeleteCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteCartService(ctx)
	// init req and assert value

	req := &cart.DeleteCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
