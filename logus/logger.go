package logus

import (
	"log"
)

func init() {
	Register(stdLog{})
}

// Logger interface collect engine log message
type Logger interface {
	Info(msg string)
	Debug(msg string)
	Warn(msg string)
	Error(msg string)
	Panic(msg string)
}

type stdLog struct{}

func (stdLog) Info(msg string) {
	log.Printf("Info %s\n", msg)
}

func (stdLog) Debug(msg string) {
	if InDebug {
		log.Printf("Debug %s\n", msg)
	}
}

func (stdLog) Warn(msg string) {
	log.Printf("Warn %s\n", msg)
}

func (stdLog) Error(msg string) {
	log.Printf("Error %s\n", msg)
}

func (stdLog) Panic(msg string) {
	log.Printf("Panic %s\n", msg)
	panic(msg)
}
