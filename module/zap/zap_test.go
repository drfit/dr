package zap

import (
	"github.com/unisx/dr/logus"
)

func ExampleZap() {
	defer func() {
		_ = recover()
	}()
	logus.Info("ping pong")
	logus.Debug("ping pong")
	logus.Warn("ping pong")
	logus.Error("ping pong")
	logus.Panic("ping pong")
	// Output:
}
