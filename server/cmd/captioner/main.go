package main

import (
	"captioner.com.ng/internal/captioner/server"
	"captioner.com.ng/internal/captioner/types"
	"github.com/spf13/pflag"
)

func init() {
	pflag.StringVarP(&types.Addr, "port", "p", ":4000", "Set the port you wish the app to run on")
	pflag.Parse()
}

func main() {
	server.Start(&types.Addr)
}