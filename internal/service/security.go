package service

import (
	"context"
	"net/http"

	wegocontext "github.com/agorago/wego/context"
	fw "github.com/agorago/wego/fw"
	e "github.com/agorago/stringdemoservice/internal/err"
)

// Secure - a middleware for security
func Secure(ctx context.Context, chain *fw.MiddlewareChain) context.Context {
	token, ok := wegocontext.Value(ctx, "Securetoken").(string)
	if !ok || (ok && token != "passpass") {
		ctx = wegocontext.SetError(ctx, e.HTTPError(ctx, http.StatusForbidden, e.SecurityException,
			map[string]interface{}{}))
		return ctx
	}
	ctx = chain.DoContinue(ctx)
	return ctx
}
