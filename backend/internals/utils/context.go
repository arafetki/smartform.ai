package utils

import (
	"context"
	"net/http"
	"time"
)

type contextKey string

const userContextKey = contextKey("user")

func ContextSetUser(r *http.Request, user *DummyUser) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func ContextGetUser(r *http.Request) *DummyUser {
	user, ok := r.Context().Value(userContextKey).(*DummyUser)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}

func ContextWithTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}
