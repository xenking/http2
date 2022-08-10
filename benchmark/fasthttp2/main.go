package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/xenking/http2"
	"github.com/xenking/http2/benchmark/common"
)

func main() {
	debug := flag.Bool("debug", true, "Debug mode")
	flag.Parse()

	cert, priv, err := common.GenerateTestCertificate("localhost:8443")
	if err != nil {
		log.Fatalln(err)
	}

	s := &fasthttp.Server{
		ReadTimeout: time.Second * 3,
		Handler:     requestHandler,
		Name:        "http2 test",
	}
	err = s.AppendCertEmbed(cert, priv)
	if err != nil {
		log.Fatalln(err)
	}

	http2.ConfigureServer(s, http2.ServerConfig{
		Debug: *debug,
	})

	err = s.ListenAndServeTLS(":8443", "", "")
	if err != nil {
		log.Fatalln(err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	if ctx.Request.Header.IsPost() {
		fmt.Fprintf(ctx, "%s\n", ctx.Request.Body())
		return
	}

	fmt.Fprintf(ctx, "Hello 21th century!\n")
}
