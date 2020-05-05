package main

import (
	"github.com/agorago/wego/cmd"
	"github.com/agorago/wego/fw"
)

func main() {
	rs := fw.MakeRegistrationService()
	cmd.SwaggerMain(rs)
}
