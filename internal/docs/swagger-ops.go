
package docs

import (
"github.com/agorago/stringdemoapi/api"
)

// swagger:route POST /uppercase Uppercase-tag UppercaseRequestWrapper
//  Uppercase - Converts the input string into upper case
// responses:
//   200: UppercaseResponseWrapper

//  UppercaseResponse - the  Uppercase service response
// swagger:response UppercaseResponseWrapper
type UppercaseResponseWrapper struct {
// in:body
Body api.UppercaseResponse
}

// swagger:parameters UppercaseRequestWrapper
type UppercaseRequestWrapper struct{
//  UppercaseRequest - the payload for Uppercase service
// in:body
Body api.UppercaseRequest
}

// swagger:route POST /count Count-tag CountRequestWrapper
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
type CountRequestWrapper struct{
//  CountRequest - the payload for Count service
// in:body
Body api.CountRequest
}

// swagger:route GET /add-numbers AddNumbers-tag AddNumbersRequestWrapper
//  AddNumbers - adds two numbers and returns the result This method illustrates a GET method implementation in WeGO since there is no request payload required
// responses:
//   200: AddNumbersResponseWrapper

//  AddNumbersResponse - the  AddNumbers service response
// swagger:response AddNumbersResponseWrapper
type AddNumbersResponseWrapper struct {
// in:body
Body api.AddNumbersResponse
}

// swagger:parameters AddNumbersRequestWrapper
type AddNumbersRequestWrapper struct{
// 
// in:body

}



