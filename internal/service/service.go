package service

import (
	"context"
	api "github.com/agorago/stringdemoapi/api"
	"strings"
)

type stringdemo struct{}

func MakeStringdemoService() stringdemo {
	return stringdemo{}
}

func (stringdemo) Uppercase(ctx context.Context, ucr *api.UpperCaseRequest) (api.UpperCaseResponse, error) {
	return api.UpperCaseResponse{V: strings.ToUpper(ucr.S)}, nil
}

func (stringdemo) Count(ctx context.Context, cr *api.CountRequest) (api.CountResponse, error) {
	return api.CountResponse{V: len(cr.S)}, nil
}

func (stringdemo) AddNumbers(ctx context.Context, arg1 int, arg2 int) (api.AddNumbersResponse, error) {
	return api.AddNumbersResponse{Sum: arg1 + arg2}, nil
}
