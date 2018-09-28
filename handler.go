package bsmratelimit_fasthttp

import (
	"github.com/bsm/ratelimit"
	"github.com/valyala/fasthttp"
)

// LimitHandler is a very simple FastHTTP adapter for `github.com/bsm/ratelimit`. It wraps original request handler with a rate-limited one
func LimitHandler(handler fasthttp.RequestHandler, limiter *ratelimit.RateLimiter, onLimitReachedHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		shouldLimit := limiter.Limit()

		if shouldLimit {
			onLimitReachedHandler(ctx)
			return
		}

		handler(ctx)
	}
}

// SimpleLimitHandler covers the most simple rate limiting case when the server only needs to return a predefined status code.
func SimpleLimitHandler(handler fasthttp.RequestHandler, limiter *ratelimit.RateLimiter, statusCode int) fasthttp.RequestHandler {
	return LimitHandler(handler, limiter, func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(statusCode)
		ctx.SetBody([]byte("You have reached maximum request limit."))
	})
}
