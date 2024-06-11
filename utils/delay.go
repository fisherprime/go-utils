// SPDX-License-Identifier: MIT
package utils

import (
	"fmt"
	"time"
)

// DelaySequencer defines a type to generate a delay time.Duration sequence.
type DelaySequencer struct {
	cache    map[uint]uint
	counter  uint
	attempts int
	infinite bool
}

const counterWrapAround = 11

// NewDelaySequencer instantiates a DelayTimer.
func NewDelaySequencer(attempts int) *DelaySequencer {
	return &DelaySequencer{
		cache:    make(map[uint]uint),
		counter:  1,
		attempts: attempts,
		infinite: attempts == -1,
	}
}

// GetDelay generates a fibonacci guided delay time.Duration.
//
// ok indicates the sequence has not terminated.
func (d *DelaySequencer) GetDelay() (dur time.Duration, ok bool) {
	switch {
	case d.infinite && int(d.counter) > counterWrapAround:
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
func fibonacci(cache map[uint]uint, n uint) uint {
	if n < 2 {
		return n
	}

	if seq, ok := cache[n]; ok {
		return seq
	}

	seq := fibonacci(cache, n-1) + fibonacci(cache, n-2)
	cache[n] = seq

	return seq
}
