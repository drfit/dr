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
	log.Printf("Info Dr. %s\n", msg)
}

func (stdLog) Debug(msg string) {
	if InDebug {
		log.Printf("Debug Dr. %s\n", msg)
	}
}

func (stdLog) Warn(msg string) {
	log.Printf("Warn Dr. %s\n", msg)
}

func (stdLog) Error(msg string) {
	log.Printf("Error Dr. %s\n", msg)
}

func (stdLog) Panic(msg string) {
	log.Printf("Panic Dr. %s\n", msg)
	panic(msg)
}
