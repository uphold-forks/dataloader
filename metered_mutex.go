package dataloader

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

// MeteredMutex wraps a sync.Mutex and tracks the time spent locked using OpenTelemetry metrics.
type MeteredMutex struct {
	mu        sync.Locker
	lockTime  time.Time
	meter     metric.Meter
	lockUsage metric.Float64Histogram // Histogram to record lock durations
}

// NewMeteredMutex creates a new MeteredMutex with the provided OpenTelemetry meter.
func NewMeteredMutex(mu sync.Locker) *MeteredMutex {
	meter := otel.GetMeterProvider().Meter("cache")
	lockUsage, err := meter.Float64Histogram(
		"mutex.lock.duration",
		metric.WithDescription("Time spent holding the mutex lock in seconds"),
	)

	if err != nil {
		panic("failed to create metered mutex")
	}

	return &MeteredMutex{
		mu:        mu,
		meter:     meter,
		lockUsage: lockUsage,
	}
}

// Lock locks the mutex and starts tracking time.
func (m *MeteredMutex) Lock() {
	m.mu.Lock()
	m.lockTime = time.Now()
}

// Unlock unlocks the mutex and records the time spent locked.
func (m *MeteredMutex) Unlock() {
	lockedDuration := time.Since(m.lockTime).Seconds()
	m.lockUsage.Record(context.Background(), lockedDuration)
	m.mu.Unlock()
}
