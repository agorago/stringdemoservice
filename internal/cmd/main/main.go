package main

import (
	"context"
	"gitlab.intelligentb.com/devops/bplus/cmd"
	"gitlab.intelligentb.com/devops/bplus/log"
	_ "gitlab.intelligentb.com/examples/stringdemo/stringdemoservice"
)

var Version = "development"
func main() {
	log.Infof(context.TODO(),"Version is %s",Version)
	cmd.Serve()
}
