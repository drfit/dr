package dr

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	ctx := NewContext()
	ctx = context.WithValue(ctx, "ping", "pong")
	pong, _ := ctx.Value("ping").(string)
	assert.Equal(t, pong, "pong")
}
