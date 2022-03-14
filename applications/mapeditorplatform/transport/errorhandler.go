package transport

import (
	"context"
	"github.com/Fighting2520/kitgo/common/errorx"
	"github.com/pkg/errors"
)

type defaultErrorHandler struct {
}

func (h *defaultErrorHandler) Handle(ctx context.Context, err error) {
	if errors.Is(err, &errorx.CodeError{}) {

	}
}

func newDefaultErrorHandler() *defaultErrorHandler {
	return &defaultErrorHandler{}
}
