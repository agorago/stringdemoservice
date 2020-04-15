package register

import (
	_ "gitlab.intelligentb.com/devops/bplus" // initialize BPlus first to make sure
	bplus "gitlab.intelligentb.com/devops/bplus/fw"
	// that all BPLUS modules are loaded
	apiregister "gitlab.intelligentb.com/examples/stringdemo/stringdemoapi/register"
	service "gitlab.intelligentb.com/examples/stringdemo/stringdemoservice/internal/service"
)

func init() {
	var sd = apiregister.GetServiceDescriptor()
	sd.ServiceToInvoke = service.MakeStringdemoService()
	for i := range sd.Operations{
		if sd.Operations[i].Name == "Uppercase"{
			sd.Operations[i].OpMiddleware = []bplus.Middleware{service.Secure,}
		}
	}
	bplus.RegisterService("stringdemo", sd)
}
