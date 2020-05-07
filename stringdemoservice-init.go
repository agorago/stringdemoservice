package stringdemoservice

import (
	"github.com/agorago/stringdemoservice/internal/register"
	"github.com/agorago/wego"
	"github.com/agorago/wego/fw"
)

const (
	StringDemoService = "StringDemoService"
)

func MakeStringDemoServiceInitializer()fw.Initializer{
	return stringDemoServiceInitializer{}
}

type stringDemoServiceInitializer struct{}

func (stringDemoServiceInitializer)ModuleName() string{
	return StringDemoService
}
// The stringdemoservice initializer
func (stringDemoServiceInitializer)Initialize(commandCatalog fw.CommandCatalog)(fw.CommandCatalog,error){
	rs,err := wego.GetWego(commandCatalog)
	if err != nil {
		return commandCatalog,err
	}
	register.RegisterStringdemoService(rs)
	return commandCatalog,nil
}
