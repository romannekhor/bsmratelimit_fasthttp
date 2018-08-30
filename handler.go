package bsmratelimit_fasthttp

import (
	"github.com/bsm/ratelimit"
	"github.com/valyala/fasthttp"
)

// LimitHandler is a very simple FastHTTP adapter for `github.com/bsm/ratelimit`. It wraps original request handler with a rate-limited one
func LimitHandler(handler fasthttp.RequestHandler, limiter *ratelimit.RateLimiter, statusCode int) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		shouldLimit := limiter.Limit()

		if shouldLimit {
			ctx.SetStatusCode(statusCode)
			ctx.SetBody([]byte("You have reached maximum request limit."))
			return
		}

		handler(ctx)
	}
}
