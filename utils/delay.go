// SPDX-License-Identifier: MIT
package utils

import (
	"fmt"
	"time"
)

// DelaySequencer defines a type to generate a delay time.Duration sequence.
type (
	DelaySequencer struct {
		cache             map[uint]uint // Cache for fibonacci numbers
		attempts          int           // Maximum number of attempts before resetting the DelaySequencer, -1 for infinite
		counterWrapAround int

		counter uint // Current counter in the Fibonacci sequence

		infinite bool // Flag indicating if there is no maximum limit on attempts
	}

	DelayOption func(d *DelaySequencer)
)

const (
	defaultAttempts          = 5
	defaultCounterWrapAround = 11
)

// NewDelaySequencer instantiates a DelayTimer.
func NewDelaySequencer(opts ...DelayOption) *DelaySequencer {
	d := &DelaySequencer{
		cache:             make(map[uint]uint),
		counter:           1,
		attempts:          defaultAttempts,
		counterWrapAround: defaultCounterWrapAround,
	}

	for _, opt := range opts {
		opt(d)
	}

	d.infinite = d.attempts == -1

	return d
}

// WithDelayAttempts configures the DelaySequencer delay attempts.
func WithDelayAttempts(attempts int) DelayOption {
	return func(d *DelaySequencer) { d.attempts = attempts }
}

// WithDelayCounterWrapAround configures the maximum fibonacci sequence to use for the delay operation.
func WithDelayCounterWrapAround(wrapAround int) DelayOption {
	return func(d *DelaySequencer) { d.counterWrapAround = wrapAround }
}

// GetDelay generates a fibonacci guided delay time.Duration.
//
// ok indicates the sequence has not terminated.
func (d *DelaySequencer) GetDelay() (dur time.Duration, ok bool) {
	switch {
	case d.infinite && int(d.counter) > d.counterWrapAround:
		// Infinite attempts.
		d.Reset()
	case !d.infinite && int(d.counter) > d.attempts:
		d.Reset()
		return
	default:
	}

	dur, _ = time.ParseDuration(fmt.Sprint(fibonacci(d.cache, d.counter), "s"))
	ok = true

	d.counter++

	return
}

// Delay delays program execution.
func (d *DelaySequencer) Delay() (proceed bool) {
	dur, proceed := d.GetDelay()
	if !proceed {
		return
	}

	timer := time.NewTimer(dur)
	<-timer.C

	return
}

// Reset the DelaySequencer counter.
func (d *DelaySequencer) Reset() { d.counter = 1 }

// fibonacci computes the nth fibonacci sequence.
func fibonacci(fibCache map[uint]uint, n uint) uint {
	if n < 2 {
		return n
	}

	if seq, ok := fibCache[n]; ok {
		return seq
	}

	seq := fibonacci(fibCache, n-1) + fibonacci(fibCache, n-2)
	fibCache[n] = seq

	return seq
}
