package main

import (
	"expvar"
	"github.com/bsm/ratelimit"
	"github.com/sviterok/bsmratelimit_fasthttp"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	// Number of successful requests
	successfulReqs = expvar.NewInt("successfulReqs")

	// Number of rate-limited requests
	rateLimitedReqs = expvar.NewInt("rateLimitedReqs")
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/hello":
			helloHandler(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	// Create a limiter struct.
	limiter := ratelimit.New(5000, time.Second)
	fasthttp.ListenAndServe(":4444", bsmratelimit_fasthttp.LimitHandler(requestHandler, limiter, onRateLimit))

}

func helloHandler(ctx *fasthttp.RequestCtx) {
	// Increment the number of successful requests
	successfulReqs.Add(1)

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Hello, World!"))
}

func onRateLimit(ctx *fasthttp.RequestCtx) {
	// Increment the number of rate-limited requests
	rateLimitedReqs.Add(1)

	ctx.SetStatusCode(fasthttp.StatusTooManyRequests)
	ctx.SetBody([]byte("Too many requests!!!"))
}
