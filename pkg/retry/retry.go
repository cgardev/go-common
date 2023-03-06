package retry

import (
	"time"
)

type RetryInput[T any] struct {
	Attempts int
	Sleep    time.Duration
	Fn       func() (T, error)
}

// Retry is a helper function that will retry a function until it succeeds or the number of attempts
// is reached.
func Retry[T any](in RetryInput[T]) (result T, err error) {
	for i := 0; i < in.Attempts; i++ {
		if i > 0 {
			in.Sleep *= 2
		}
		result, err = in.Fn()
		if err == nil {
			return result, nil
		}
	}

	return result, ErrAllAttemptsFailed.New(
		"after %d attempts", in.Attempts,
	).WithUnderlyingErrors(err)
}
