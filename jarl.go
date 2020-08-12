package Jarl

import (
	"fmt"
	"strings"
	"time"
)

const (
	Infinite  = -1
	Immediate = 1 * time.Nanosecond
	maxErrors = 512
)

type RetryFunction func() error
type OnRetryFunction func(error)

type Retry struct {
	fn         RetryFunction
	retryFn    OnRetryFunction
	maxRetries int
	count      int
	delay      time.Duration
	errors     []string
}

func NewRetry(fn RetryFunction, maxRetries int, delay time.Duration) *Retry {
	return &Retry{
		fn:         fn,
		maxRetries: maxRetries,
		count:      0,
		delay:      delay,
	}
}

func (r *Retry) Run() error {
	for {
		err := r.fn()
		if err != nil {
			if r.retryFn != nil {
				r.retryFn(err)
			}

			if r.maxRetries > 0 && r.count < maxErrors {
				r.errors = append(r.errors, fmt.Sprintf("#%d : %s", r.count, err.Error()))
			}
		}

		if r.maxRetries > 0 {
			if r.count == r.maxRetries {
				break
			}
			r.count++
		}

		time.Sleep(r.delay)
	}

	return fmt.Errorf("Failed attempts:\n%s", strings.Join(r.errors, "\n"))
}
