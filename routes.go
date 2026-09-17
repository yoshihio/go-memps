package main

import "github.com/valyala/fasthttp"

func routes(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		ctx.SetBodyString("Hello, World!")
	default:
		ctx.SetBodyString("Not Found")
	}
}
