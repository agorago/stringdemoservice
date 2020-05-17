package register

import (
	fw "github.com/agorago/wego/fw"
	// that all WeGO modules are loaded
	apiregister "github.com/agorago/stringdemoapi/register"
)



func RegisterStringdemoService(rs fw.RegistrationService,
		serviceToInvoke interface{},
		middlewares []fw.Middleware,
		proxyMiddlewares []fw.Middleware) {
	var sd = apiregister.GetServiceDescriptor(proxyMiddlewares)
	sd.ServiceToInvoke = serviceToInvoke
	for i := range sd.Operations {
		if sd.Operations[i].Name == "Uppercase" {
			sd.Operations[i].OpMiddleware = middlewares
		}
	}
	rs.RegisterService("stringdemo", sd)
}
