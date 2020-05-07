package test

import (
	"github.com/agorago/stringdemoapi"
	"github.com/agorago/stringdemoservice"
	"github.com/agorago/wego"
	_ "github.com/agorago/wego" // to initialize the http module
	wegotest "github.com/agorago/wego/test"
	"testing"
)

func TestMain(m *testing.M) {
	wegotest.BDD(m, FeatureContext, wego.MakeWegoInitializer(),
		stringdemoapi.MakeStringDemoApiInitializer(),
		stringdemoservice.MakeStringDemoServiceInitializer(),)
}
