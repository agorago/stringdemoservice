package main

import (
	"context"
	"github.com/agorago/stringdemoservice/internal/register"
	"github.com/agorago/wego/cmd"
	"github.com/agorago/wego/fw"
	"github.com/agorago/wego/log"
)

var Version = "development"

func main() {
	log.Infof(context.TODO(), "Version is %s", Version)
	rs := fw.MakeRegistrationService()
	register.RegisterStringdemoService(rs)
	cmd.Serve()
}
