package main

import (
	"context"
	"github.com/agorago/stringdemoapi"
	"github.com/agorago/stringdemoservice"
	"github.com/agorago/wego"
	"github.com/agorago/wego/cmd"
	"github.com/agorago/wego/log"
)

var Version = "development"

func main() {
	log.Infof(context.TODO(), "Version is %s", Version)

	cmd.Serve(wego.MakeWegoInitializer(),
		stringdemoapi.MakeStringDemoApiInitializer(),
		stringdemoservice.MakeStringDemoServiceInitializer(),)
}
