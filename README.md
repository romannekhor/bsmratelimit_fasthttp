# bsm/ratelimit for FastHTTP

## What is it?
This is a very simple FastHTTP adapter for `github.com/bsm/ratelimit` library.

## How to use it?

Example:

```golang
package main

import (
	"github.com/valyala/fasthttp"
	"github.com/bsm/ratelimit"
	"time"
	"github.com/romannekhor/bsmratelimit_fasthttp"
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

```
