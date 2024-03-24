package main

import (
	"fmt"

	"captioner.com.ng/api/server"
	"captioner.com.ng/config"
	"github.com/spf13/pflag"
)

var values config.Config = *&config.Config{}

func init() {
	values.Version = 0.1
	pflag.StringVarP(&values.Addr, "addr", "a", ":4000", "Set the port you wish the app to run on")
	pflag.StringVarP(&values.MongoURI, "mongo", "m", "mongodb://localhost:27017/captioner", "Set MongoDB URI")
	pflag.Float32P("version", "v", values.Version, "Application Version")
	pflag.Parse()
}

func main() {
	fmt.Println()
	server.Start(&values.Addr, &values.MongoURI)
}
