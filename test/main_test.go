package test

import (
	"github.com/agorago/stringdemoapi/api"
	"github.com/agorago/stringdemoapi/proxy"
	"github.com/agorago/stringdemoservice/internal/register"
	_ "github.com/agorago/wego" // to initialize the http module
	"github.com/agorago/wego/fw"
	wegohttp "github.com/agorago/wego/http"
	wegotest "github.com/agorago/wego/test"
	"testing"
)
var stringdemoProxy api.StringDemoService
func TestMain(m *testing.M) {
    // make sure that the registration service is initialized with Wego
	rs := fw.MakeRegistrationService()
	register.RegisterStringdemoService(rs)
	// Initialize the wego proxy
	wegoProxy := wegohttp.MakeProxyService(rs)
	stringdemoProxy = proxy.MakeStringdemoProxy(wegoProxy)
	wegotest.BDD(m, FeatureContext)
}
