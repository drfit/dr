package dr

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestStdLog(t *testing.T) {
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	for _, test := range []struct {
		input  string
		output string
		logFun func(string)
	}{
		{
			input:  "ping pong",
			output: ".* Info Dr. ping pong",
			logFun: logus.Info,
		},
		{
			input:  "ping pong",
			output: ".* Warn Dr. ping pong",
			logFun: logus.Warn,
		},
		{
			input:  "ping pong",
			output: ".* Error Dr. ping pong",
			logFun: logus.Error,
		},
	} {
		buf.Reset()
		test.logFun(test.input)
		assert.Regexp(t, test.output, buf.String())
	}
}

func TestStdLog_Debug(t *testing.T) {
	oldInDebug := InDebug

	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	InDebug = false
	logus.Debug("ping pong")
	assert.Equal(t, "", buf.String())

	InDebug = true
	buf.Reset()
	logus.Debug("ping pong")
	assert.Regexp(t, ".* Debug Dr. ping pong", buf.String())

	InDebug = oldInDebug
}

func TestStdLog_Panic(t *testing.T) {
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	defer func() {
		assert.Regexp(t, ".* Panic Dr. ping pong", buf.String())
		err := recover()
		msg, _ := err.(string)
		assert.Equal(t, "ping pong", msg)
	}()

	logus.Panic("ping pong")
}
