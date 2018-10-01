package main

import (
	"github.com/bsm/ratelimit"
	"github.com/sviterok/bsmratelimit_fasthttp"
	"github.com/valyala/fasthttp"
	"time"
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
	statusCode := fasthttp.StatusTooManyRequests
	fasthttp.ListenAndServe(":4444", bsmratelimit_fasthttp.SimpleLimitHandler(requestHandler, limiter, statusCode))

}

func helloHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Hello, World!"))
}
