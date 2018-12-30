package dr

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ctxKey string

func TestContextWithValue(t *testing.T) {
	ctx := NewContext()
	ctx = context.WithValue(ctx, ctxKey("ping"), "pong")
	pong, _ := ctx.Value(ctxKey("ping")).(string)
	assert.Equal(t, pong, "pong")
}
