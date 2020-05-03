package test

import (
	"context"
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/agorago/stringdemoapi/api"
	"github.com/agorago/stringdemoapi/proxy"
	_ "github.com/agorago/stringdemoservice" // invoke init() on hello functions
	bplusc "github.com/agorago/wego/context"
	bpluse "github.com/agorago/wego/err"
)

type stringdemoTestStruct struct {
	Ur  api.UpperCaseResponse
	Cr  api.CountResponse
	Anr api.AddNumbersResponse
	Err error
}

// FeatureContext - the GODOG feature context that defines the step definitions for the BDD tests
// that need to be executed for hello service
func FeatureContext(s *godog.Suite) {
	s.Step(`^I invoke Uppercase with "([^"]*)" and a secure token with value "([^"]*)"$`, hts.iInvokeUppercaseWithSecureToken)
	s.Step(`^I must get back an upper case return value of "([^"]*)"$`, hts.iMustGetBackAnUpperCaseReturnValueOf)
	s.Step(`^I invoke Count with "([^"]*)"$`, hts.iInvokeCountWith)
	s.Step(`^I must get back a count return value of (\d+)$`, hts.iMustGetBackACountReturnValueOf)
	s.Step(`^I invoke AddNumbers with (\d+) and (\d+)$`, hts.iInvokeAddNumbersWithAnd)
	s.Step(`^I must get back an addnumbers return value of (\d+)$`, hts.iMustGetBackAnAddnumbersReturnValueOf)
	s.Step(`^I invoke Uppercase with "([^"]*)" without secure token$`, hts.iInvokeUppercaseWithoutSecureToken)
	s.Step(`^I must get back an error with http error code (\d+)$`, hts.iMustGetBackAnErrorWithHTTPErrorCode)
}

var hts = &stringdemoTestStruct{}
var stringdemoProxy = proxy.Stringdemo{}

func (sts *stringdemoTestStruct) iInvokeUppercaseWithSecureToken(arg string, token string) error {
	uc := api.UpperCaseRequest{S: arg}
	ctx := context.TODO()
	ctx = bplusc.Add(ctx, "SecureToken", token)

	resp, err := stringdemoProxy.Uppercase(ctx, &uc)
	if err != nil {
		sts.Err = err
		return nil
	}
	sts.Ur = resp
	return nil
}

func (sts *stringdemoTestStruct) iInvokeUppercaseWithoutSecureToken(arg string) error {
	uc := api.UpperCaseRequest{S: arg}
	ctx := context.TODO()

	_, err := stringdemoProxy.Uppercase(ctx, &uc)
	if err == nil {
		return fmt.Errorf("The Uppercase must have errored out if secure token is not passed")
	}
	sts.Err = err
	return nil
}

func (sts *stringdemoTestStruct) iMustGetBackAnErrorWithHTTPErrorCode(arg1 int) error {
	err, ok := sts.Err.(bpluse.BPlusError)
	if !ok {
		return fmt.Errorf("Error is not of type BplusError. it is %#v\n", err)
	}
	if arg1 != err.HTTPErrorCode {
		return fmt.Errorf("error codes dont match. Expected = %d.Actual = %d. Full error is %#v",
			arg1, err.HTTPErrorCode, err)
	}
	return nil
}

func (sts *stringdemoTestStruct) iMustGetBackAnUpperCaseReturnValueOf(arg1 string) error {
	if arg1 != sts.Ur.V {
		return fmt.Errorf("test failure: expected = %s. actual = %s", arg1, sts.Ur.V)
	}
	return nil
}

func (sts *stringdemoTestStruct) iInvokeCountWith(arg string) error {
	uc := api.CountRequest{S: arg}
	resp, err := stringdemoProxy.Count(context.TODO(), &uc)
	if err != nil {
		return fmt.Errorf("cannot invoke the hello proxy. err = %s", err.Error())
	}
	sts.Cr = resp
	return nil
}

func (sts *stringdemoTestStruct) iMustGetBackACountReturnValueOf(arg1 int) error {
	if arg1 != sts.Cr.V {
		return fmt.Errorf("test failure: expected = %d. actual = %d", arg1, sts.Cr.V)
	}
	return nil
}

func (sts *stringdemoTestStruct) iInvokeAddNumbersWithAnd(arg1, arg2 int) error {
	resp, err := stringdemoProxy.AddNumbers(context.TODO(), arg1, arg2)
	if err != nil {
		return fmt.Errorf("cannot invoke the hello proxy. err = %s", err.Error())
	}
	sts.Anr = resp
	return nil
}

func (sts *stringdemoTestStruct) iMustGetBackAnAddnumbersReturnValueOf(arg1 int) error {
	if arg1 != sts.Anr.Sum {
		return fmt.Errorf("test failure: expected = %d. actual = %d", arg1, sts.Anr.Sum)
	}
	return nil
}
