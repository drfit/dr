package dr

import "context"

// Context Dr. main entry
type Context interface {
	context.Context
}

// NewContext instance a new Context object
func NewContext() Context {
	return context.Background()
}
