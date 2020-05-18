package service

import (
	"context"
	api "github.com/agorago/stringdemoapi/api"
	"strings"
)

type stringdemo struct{
}

func MakeStringdemoService() stringdemo {
	return stringdemo{}
}

func (stringdemo) Uppercase(_ context.Context, ucr *api.UppercaseRequest) (api.UppercaseResponse, error) {
	return api.UppercaseResponse{V: strings.ToUpper(ucr.S)}, nil
}

func (stringdemo) Count(_ context.Context, cr *api.CountRequest) (api.CountResponse, error) {
	return api.CountResponse{V: len(cr.S)}, nil
}

func (stringdemo) AddNumbers(_ context.Context, arg1 int, arg2 int) (api.AddNumbersResponse, error) {
	return api.AddNumbersResponse{Sum: arg1 + arg2}, nil
}
