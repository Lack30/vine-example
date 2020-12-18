package main

import (
	ccli "github.com/lack-io/cli"

	goproto "github.com/lack-io/vineadm/gogenerator/examples/goproto-gen"
)

func main() {
	g := goproto.New()
	g.BindFlags(ccli.CommandLine)
	ccli.CommandLine.RunAndExitOnError()
	goproto.Run(g)
}