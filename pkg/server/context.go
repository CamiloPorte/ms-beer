package server

import (
	"context"
)

var (
	contextKeyServerID        = contextKey("id")
	contextKeyXForwardedFor   = contextKey("xForwardedFor")
	contextKeyXForwardedProto = contextKey("xForwardedProto")
	contextKeyEndpoint        = contextKey("endpoint")
	contextKeyClientIP        = contextKey("clientIP")
)

type contextKey string

func (c contextKey) String() string {
	return "server" + string(c)
}

// ID gets the name server from context
func ID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(contextKeyServerID).(string)
	return id, ok
}
