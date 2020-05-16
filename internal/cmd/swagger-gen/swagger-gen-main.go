package main

import (
	"github.com/agorago/stringdemoservice"
	"github.com/agorago/wego/cmd"
)

func main() {
	cmd.SwaggerMain(stringdemoservice.Initializers...)
}
