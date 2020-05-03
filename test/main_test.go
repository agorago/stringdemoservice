package test

import (
	"testing"

	bplustest "github.com/agorago/wego/test"
)

func TestMain(m *testing.M) {
	bplustest.BDD(m, FeatureContext)
}
