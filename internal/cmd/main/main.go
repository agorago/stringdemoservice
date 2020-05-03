package main

import (
	"context"
	"github.com/agorago/wego/cmd"
	"github.com/agorago/wego/log"
	_ "github.com/agorago/stringdemoservice"
)

var Version = "development"

func main() {
	log.Infof(context.TODO(), "Version is %s", Version)
	cmd.Serve()
}
