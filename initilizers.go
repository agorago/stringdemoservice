package stringdemoservice

import (
	"github.com/agorago/stringdemoapi"
	"github.com/agorago/wego"
	"github.com/agorago/wego/fw"
)

var Initializers = []fw.Initializer{
	wego.MakeWegoInitializer(),
	stringdemoapi.MakeStringDemoApiInitializer(),
	MakeStringDemoServiceInitializer(),
}
