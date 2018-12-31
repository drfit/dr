package zap

import (
	"github.com/unisx/dr/logus"
	zap "github.com/unisx/logus"
)

func init() {
	logus.Register(drZap{})
}

type drZap struct{}

const (
	logDr = "Dr.Log"
	msgDr = "Dr.Msg"
)

func (drZap) Info(msg string) {
	zap.Info(logDr, zap.String(msgDr, msg))
}

func (drZap) Debug(msg string) {
	zap.Debug(logDr, zap.String(msgDr, msg))
}

func (drZap) Warn(msg string) {
	zap.Warn(logDr, zap.String(msgDr, msg))
}

func (drZap) Error(msg string) {
	zap.Error(logDr, zap.String(msgDr, msg))
}

func (drZap) Panic(msg string) {
	zap.Panic(logDr, zap.String(msgDr, msg))
}
