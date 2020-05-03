package register

import (
	_ "github.com/agorago/wego" // initialize BPlus first to make sure
	bplus "github.com/agorago/wego/fw"
	// that all BPLUS modules are loaded
	apiregister "github.com/agorago/stringdemoapi/register"
	service "github.com/agorago/stringdemoservice/internal/service"
)

func init() {
	var sd = apiregister.GetServiceDescriptor()
	sd.ServiceToInvoke = service.MakeStringdemoService()
	for i := range sd.Operations {
		if sd.Operations[i].Name == "Uppercase" {
			sd.Operations[i].OpMiddleware = []bplus.Middleware{service.Secure}
		}
	}
	bplus.RegisterService("stringdemo", sd)
}
