package dataloader

import (
	"context"
)

type TraceLoadFinishFunc[V any] func(Thunk[V])
type TraceLoadManyFinishFunc[V any] func(ThunkMany[V])
type TraceBatchFinishFunc[V any] func([]*Result[V])
type TraceWaitFinishFunc[V any] func(keys []V)

// Tracer is an interface that may be used to implement tracing.
type Tracer[K comparable, V any] interface {
	// TraceLoad will trace the calls to Load.
	TraceLoad(ctx context.Context, key K) (context.Context, TraceLoadFinishFunc[V])
	// TraceLoadMany will trace the calls to LoadMany.
	TraceLoadMany(ctx context.Context, keys []K) (context.Context, TraceLoadManyFinishFunc[V])
	// TraceBatch will trace data loader batches.
	TraceBatch(ctx context.Context, keys []K) (context.Context, TraceBatchFinishFunc[V])
	// TraceWait will trace waits between load and batches.
	TraceWait(ctx context.Context) (context.Context, TraceWaitFinishFunc[K])
}

// NoopTracer is the default (noop) tracer
type NoopTracer[K comparable, V any] struct{}

// TraceLoad is a noop function
func (NoopTracer[K, V]) TraceLoad(ctx context.Context, key K) (context.Context, TraceLoadFinishFunc[V]) {
	return ctx, func(Thunk[V]) {}
}

// TraceLoadMany is a noop function
func (NoopTracer[K, V]) TraceLoadMany(ctx context.Context, keys []K) (context.Context, TraceLoadManyFinishFunc[V]) {
	return ctx, func(ThunkMany[V]) {}
}

// TraceBatch is a noop function
func (NoopTracer[K, V]) TraceBatch(ctx context.Context, keys []K) (context.Context, TraceBatchFinishFunc[V]) {
	return ctx, func(result []*Result[V]) {}
}

// TraceWait is a noop function
func (NoopTracer[K, V]) TraceWait(ctx context.Context) (context.Context, TraceWaitFinishFunc[K]) {
	return ctx, func([]K) {}
}
