// Package node provides an ultra-lightweight HTTP server for serving a single checkbox image.
//go:generate go-bindata -nomemcopy -nometadata img/
package main

import (
	"bytes"
	"log"

	util "github.com/kayteh/reachability"
	"github.com/valyala/fasthttp"
)

func main() {
	log.Println("starting...")
	err := fasthttp.ListenAndServe(util.EnvDef("ADDR", "0.0.0.0:80"), handleWrapper())
	if err != nil {
		log.Fatal(err)
	}
}

func handleWrapper() fasthttp.RequestHandler {
	okPath := []byte("/ok.png")
	okFile, err := imgOkPngBytes()
	if err != nil {
		log.Fatalln("failed to preload image: ", err)
	}
	return func(ctx *fasthttp.RequestCtx) {
		if bytes.Equal(ctx.Path(), okPath) {
			ctx.SetStatusCode(200)
			ctx.SetContentType("image/png")
			_, err := ctx.Write(okFile)
			if err != nil {
				ctx.SetStatusCode(500)
			}

			return
		}

		ctx.SetStatusCode(404)
		ctx.SetContentType("text/html")
		ctx.WriteString("<!doctype html><meta charset=utf-8><h1>This server is reachable, but there is no content.</h1>")
	}
}
