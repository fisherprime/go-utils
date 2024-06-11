package utils

import (
	"context"
	"fmt"
)

// MonitorChannels `error`s & completion status.
//
// errPrefix should be in the singular form.
//
// Usage:
//
//		done, errChan := make(chan bool), make(chan error, defaultChanBufferSize)
//		numGoroutines := 5
//
//	 wg := sync.WaitGroup{}
//	 wg.Add(numGoroutines)
//
//		for index :=0; index < numGoroutines; index++ {
//			go func() {
//				if err := someFunc(); err != nil {errChan <- err}
//				wg.Done()
//			}()
//		}
//
//		go func() {
//			wg.Wait()
//
//			close(done)
//			close(errChan)
//		}()
//
//		if err := utils.MonitorChannels(ctx, numGoroutines, done, errChan, "error prefix"); err != nil {
//			return err
//		}
func MonitorChannels(ctx context.Context, operations int, done chan bool, errChan chan error, errPrefix string) (err error) {
	if operations < 1 {
		err = fmt.Errorf("%s %w: %d", errPrefix, ErrInvalidGoroutineCount, operations)
		return
	}

	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:
		var routinesTerminated bool
		for index := 0; index < operations; index++ {
			select {
			case _, routinesTerminated = <-done:
				// On channel close.
				routinesTerminated = !routinesTerminated
			case e, proceed := <-errChan:
				if !proceed {
					break
				}

				if err != nil {
					err = fmt.Errorf("%v, %w", err, e)
				} else {
					err = fmt.Errorf("%s %w", errPrefix, e)
				}
			}

			if routinesTerminated {
				break
			}
		}
	}

	return
}
