package service

import (
	"context"
	"net/http"

	bplusc "gitlab.intelligentb.com/devops/bplus/context"
	bplus "gitlab.intelligentb.com/devops/bplus/fw"
	e "gitlab.intelligentb.com/examples/stringdemo/stringdemoservice/internal/err"
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
