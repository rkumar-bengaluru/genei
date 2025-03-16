package zcontext

import (
	"context"
	"time"
)

type contextWithoutDeadline struct {
	ctx context.Context
}

func (l *contextWithoutDeadline) Deadline() (time.Time, bool) { return time.Time{}, false }
func (l *contextWithoutDeadline) Done() <-chan struct{}       { return nil }
func (l *contextWithoutDeadline) Err() error                  { return nil }
func (l *contextWithoutDeadline) Value(key interface{}) interface{} {
	return l.ctx.Value(key)
}

// ContextWithoutDeadline creates a copy of ctx without any deadline
func ContextWithoutDeadline(ctx context.Context) context.Context {
	return &contextWithoutDeadline{ctx}
}
