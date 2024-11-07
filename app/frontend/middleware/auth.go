package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "user_id"

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//todo
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}
