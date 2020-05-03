package docs

import (
	"github.com/agorago/stringdemoapi/api"
)

// swagger:route POST /stringdemo/uppercase Uppercase-tag UppercaseRequestWrapper
//  Uppercase - Converts the input string into upper case
// responses:
//   200: UppercaseResponseWrapper

//  UpperCaseResponse - the  Uppercase service response
// swagger:response UppercaseResponseWrapper
type UppercaseResponseWrapper struct {
	// in:body
	Body api.UpperCaseResponse
}

// swagger:parameters UppercaseRequestWrapper
type UppercaseRequestWrapper struct {
	//  UpperCaseRequest - the payload for Uppercase service
	// in:body
	Body api.UpperCaseRequest
}

// swagger:route POST /stringdemo/count Count-tag CountRequestWrapper
//  Count - returns the length of the input string
// responses:
//   200: CountResponseWrapper

//  CountResponse - the  Count service response
// swagger:response CountResponseWrapper
type CountResponseWrapper struct {
	// in:body
	Body api.CountResponse
}

// swagger:parameters CountRequestWrapper
type CountRequestWrapper struct {
	//  CountRequest - the payload for Count service
	// in:body
	Body api.CountRequest
}

// swagger:route GET /stringdemo/add-numbers AddNumbers-tag
//  AddNumbers - adds two numbers and returns the result This method illustrates a GET method implementation in BPlus since there is no request payload required
// responses:
//   200: AddNumbersResponseWrapper

//  AddNumbersResponse - the  AddNumbers service response
// swagger:response AddNumbersResponseWrapper
type AddNumbersResponseWrapper struct {
	// in:body
	Body api.AddNumbersResponse
}
