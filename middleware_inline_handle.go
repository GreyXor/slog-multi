package slogmulti

import (
	"context"
	"log/slog"
)

// NewHandleInlineMiddleware is a shortcut to a middleware that implements only the `Handle` method.
func NewHandleInlineMiddleware(handleFunc func(ctx context.Context, record slog.Record, next func(context.Context, slog.Record) error) error) Middleware {
	return func(next slog.Handler) slog.Handler {
		return &HandleInlineMiddleware{
			next:       next,
			handleFunc: handleFunc,
		}
	}
}

type HandleInlineMiddleware struct {
	next       slog.Handler
	handleFunc func(ctx context.Context, record slog.Record, next func(context.Context, slog.Record) error) error
}

// Implements slog.Handler
func (h *HandleInlineMiddleware) Enabled(ctx context.Context, level slog.Level) bool {
	return h.next.Enabled(ctx, level)
}

// Implements slog.Handler
func (h *HandleInlineMiddleware) Handle(ctx context.Context, record slog.Record) error {
	return h.handleFunc(ctx, record, h.next.Handle)
}

// Implements slog.Handler
func (h *HandleInlineMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewHandleInlineMiddleware(h.handleFunc)(h.next.WithAttrs(attrs))
}

// Implements slog.Handler
func (h *HandleInlineMiddleware) WithGroup(name string) slog.Handler {
	return NewHandleInlineMiddleware(h.handleFunc)(h.next.WithGroup(name))
}
