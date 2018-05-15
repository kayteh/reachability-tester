// Package app is a binary for building then serving the reachability test.
//go:generate go-bindata -nomemcopy templates img
package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net"

	util "github.com/kayteh/reachability"
	"github.com/valyala/fasthttp"
)

type data struct {
	SiteTitle       string
	HeaderImage     string
	GroupHeaderName string
	TargetsCount    int
	Timeout         int
	ColorScheme     colorScheme
	Regions         []struct {
		Name          string
		NetworkGroups []struct {
			Name    string
			Subnets []struct {
				Name   string
				Subnet string
				Target string
			}
		}
	}
}

type colorScheme struct {
	MainBg           string
	MainText         string
	GroupHeaderText  string
	GroupHeaderBg    string
	GroupCardBg      string
	RegionHeaderText string
	RegionHeaderBg   string
	RegionCardBg     string
	SubnetText       string
	SubnetBg         string
	FailedSubnetText string
	FailedSubnetBg   string
	SubnetBorder     string
	Border           string
}

func main() {
	err := fasthttp.ListenAndServe(util.EnvDef("ADDR", "0.0.0.0:80"), handlerWrapper())
	if err != nil {
		log.Fatalln("server failed: ", err)
	}
}

func (d data) getTargetsCount() (i int) {
	for _, region := range d.Regions {
		for _, group := range region.NetworkGroups {
			i += len(group.Subnets)
		}
	}

	return
}

func ipfilter(i string) (o string) {
	o, _, err := net.SplitHostPort(i)
	if err != nil {
		o = i
	}

	return
}

func handlerWrapper() fasthttp.RequestHandler {
	conf, err := ioutil.ReadFile(util.EnvDef("CONFIG_PATH", "/etc/reachability/config.json"))
	if err != nil {
		log.Fatalln("config failed to load: ", err)
	}

	var d data
	err = json.Unmarshal(conf, &d)
	if err != nil {
		log.Fatalln("config failed to parse: ", err)
	}

	d.TargetsCount = d.getTargetsCount()
	if d.Timeout == 0 {
		d.Timeout = 10
	}

	fm := template.FuncMap{
		"ipfilter": ipfilter,
	}
	tmpl := template.Must(template.New("main").Funcs(fm).Parse(string(util.B(templatesIndexHtmlTmplBytes()))))
	buf := bytes.Buffer{}
	tmpl.Execute(&buf, d)

	index := buf.Bytes()
	errorImg := util.B(imgErrorPngBytes())

	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			ctx.Success("text/html", index)
		case "/error.png":
			ctx.Success("image/png", errorImg)
		default:
			ctx.Error("Not Found", 404)
		}
	}
}
