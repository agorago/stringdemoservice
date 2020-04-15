package main

import (
"gitlab.intelligentb.com/devops/bplus/cmd"
_ "gitlab.intelligentb.com/examples/stringdemo/stringdemoservice"
)

func main() {
	cmd.Serve()
}
