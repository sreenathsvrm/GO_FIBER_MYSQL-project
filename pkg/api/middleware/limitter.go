package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type reqCountDetails struct {
	count            int
	firstRequestTime time.Time
}

type Limit struct {
	maxLimit          int
	duration          time.Duration
	requestCountStore map[string]*reqCountDetails
	mu                sync.Mutex
}

func LimitMiddlewarehHandler() fiber.Handler {
	limiter := &Limit{
		maxLimit:          60,
		duration:          60 * time.Second,
		requestCountStore: make(map[string]*reqCountDetails),
		mu:                sync.Mutex{},
	}

	return func(c *fiber.Ctx) error {

		clientIP := c.IP()

		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		now := time.Now()
		for ip, reqDetails := range limiter.requestCountStore {

			if now.Sub(reqDetails.firstRequestTime) >= limiter.duration {
				delete(limiter.requestCountStore, ip)
			}
		}

		reqDetails, ok := limiter.requestCountStore[clientIP]

		if ok && reqDetails.count > limiter.maxLimit {
			return fiber.ErrTooManyRequests
		}

		if !ok {
			limiter.requestCountStore[clientIP] = &reqCountDetails{
				count:            2,
				firstRequestTime: time.Now(),
			}
			return c.Next()
		}

		reqDetails.count++
		return c.Next()
	}
}
