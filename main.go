package log

import (
	"errors"
	"io"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	err = 1 << iota
	warn
	info

	infoColor  = "\033[34m"
	warnColor  = "\033[33m"
	errColor   = "\033[31m"
	strColor   = "\033[32m"
	resetColor = "\033[0m"
)

type Opts struct {
	funcName   bool
	fileName   bool
	lineNumber bool
	date       bool
	time       bool
}

type Logger struct {
	out io.Writer
	*Opts
}

type Messages []Msg

type Msg string

func New(out io.Writer, opts *Opts) (*Logger, error) {
	if opts == nil {
		return nil, errors.New("opts is nil")
	}
	return &Logger{
		out:  out,
		Opts: opts,
	}, nil
}

func Str(key, value string) Msg {
	return Msg(strColor + key + " : " + value + resetColor)
}

func Error(message error) Msg {
	if message == nil {
		message = errors.New("")
	}
	return Msg(errColor + "Error: " + message.Error() + resetColor)
}

func Warning(warning string) Msg {
	return Msg(warnColor + "Warning: " + warning + resetColor)
}

func Info(message string) Msg {
	return Msg(infoColor + "Info: " + message + resetColor)
}

func (l *Logger) Print(messages ...Msg) {
	pc, filePath, lineNumber, _ := runtime.Caller(1)
	if len(messages) < 1 {
		return
	}
	if l.date {
		l.out.Write([]byte(time.Now().Format(time.DateOnly) + " "))
	}
	if l.time {
		l.out.Write([]byte(time.Now().Format(time.TimeOnly) + " "))
	}
	if l.fileName {
		l.out.Write([]byte(filePath + " "))
	}
	if l.lineNumber {
		l.out.Write([]byte(strconv.Itoa(lineNumber) + " "))
	}
	if !l.fileName && l.lineNumber && l.funcName {
		funcName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
		l.out.Write([]byte(funcName + "() " + strconv.Itoa(lineNumber) + " "))
	}
	if len(messages) == 1 {
		l.out.Write([]byte(messages[0] + "\n"))
		return
	}
	if !l.date && !l.time && !l.fileName && !l.lineNumber && !l.funcName {
		for _, message := range messages {
			l.out.Write([]byte(message + "\n"))
		}
		return
	}
	l.out.Write([]byte(":\n"))
	for _, message := range messages {
		l.out.Write([]byte("\t\t" + message + "\n"))
	}
}
