package drL

import (
	drLog "github.com/unisx/dr/logus"
)

func ExampleDrL() {
	defer func() {
		_ = recover()
	}()
	drLog.Info("ping pong")
	drLog.Debug("ping pong")
	drLog.Warn("ping pong")
	drLog.Error("ping pong")
	drLog.Panic("ping pong")
	// Output:
}