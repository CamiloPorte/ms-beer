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

func ID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(contextKeyServerID).(string)
	return id, ok
}

func XForwardedFor(ctx context.Context) (string, bool) {
	xForwardedFor, ok := ctx.Value(contextKeyXForwardedFor).(string)
	return xForwardedFor, ok
}

func XForwardedProto(ctx context.Context) (string, bool) {
	xForwardedProto, ok := ctx.Value(contextKeyXForwardedProto).(string)
	return xForwardedProto, ok
}

func Endpoint(ctx context.Context) (string, bool) {
	endpoint, ok := ctx.Value(contextKeyEndpoint).(string)
	return endpoint, ok
}

func ClientIP(ctx context.Context) (string, bool) {
	clientIP, ok := ctx.Value(contextKeyClientIP).(string)
	return clientIP, ok
}
