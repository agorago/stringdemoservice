package register

import (
	fw "github.com/agorago/wego/fw"
	// that all WeGO modules are loaded
	apiregister "github.com/agorago/stringdemoapi/register"
	service "github.com/agorago/stringdemoservice/internal/service"
)



func RegisterStringdemoService(rs fw.RegistrationService) {
	var sd = apiregister.GetServiceDescriptor()
	sd.ServiceToInvoke = service.MakeStringdemoService()
	for i := range sd.Operations {
		if sd.Operations[i].Name == "Uppercase" {
			sd.Operations[i].OpMiddleware = []fw.Middleware{service.Secure}
		}
	}
	rs.RegisterService("stringdemo", sd)
}
