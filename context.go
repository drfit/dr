package dr

import "context"

type Context interface {
	context.Context
}

func NewContext() Context {
	return context.Background()
}
