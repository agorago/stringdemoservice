package service

import (
	"context"
	"net/http"

	bplusc "github.com/agorago/wego/context"
	bplus "github.com/agorago/wego/fw"
	e "github.com/agorago/stringdemoservice/internal/err"
)

// Secure - a middleware for security
func Secure(ctx context.Context, chain *bplus.MiddlewareChain) context.Context {
	token, ok := bplusc.Value(ctx, "Securetoken").(string)
	if !ok || (ok && token != "passpass") {
		ctx = bplusc.SetError(ctx, e.MakeBplusErrorWithErrorCode(ctx, http.StatusForbidden, e.SecurityException,
			map[string]interface{}{}))
		return ctx
	}
	ctx = chain.DoContinue(ctx)
	return ctx
}
