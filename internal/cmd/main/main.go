package main

import (
	"context"
	"github.com/agorago/stringdemoservice"
	"github.com/agorago/wego/cmd"
	"github.com/agorago/wego/log"
)

var Version = "development"

func main() {
	log.Infof(context.TODO(), "Version is %s", Version)
	cmd.Serve(stringdemoservice.Initializers...)
}
