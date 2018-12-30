package drL

import (
	drLog "github.com/unisx/dr/logus"
	"github.com/unisx/logus"
)

func init() {
	drLog.Register(drL{})
}

type drL struct{}

const (
	logDr = "Dr.Log"
	msgDr = "Dr.Msg"
)

func (drL) Info(msg string) {
	logus.Info(logDr, logus.String(msgDr, msg))
}

func (drL) Debug(msg string) {
	logus.Debug(logDr, logus.String(msgDr, msg))
}

func (drL) Warn(msg string) {
	logus.Warn(logDr, logus.String(msgDr, msg))
}

func (drL) Error(msg string) {
	logus.Error(logDr, logus.String(msgDr, msg))
}

func (drL) Panic(msg string) {
	logus.Panic(logDr, logus.String(msgDr, msg))
}
