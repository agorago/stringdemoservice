package stringdemoservice

import (
	"github.com/agorago/stringdemoapi"
	"github.com/agorago/stringdemoservice/internal/register"
	"github.com/agorago/stringdemoservice/internal/service"
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
	register.RegisterStringdemoService(rs,service.MakeStringdemoService(),
		[]fw.Middleware {service.MakeSecure()},
		stringdemoapi.GetProxyMiddlewares())
	return commandCatalog,nil
}
