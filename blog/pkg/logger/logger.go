package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// 设置日志级别
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func NewLogger(write io.Writer, prefix string, flag int) *Logger {
	l := log.New(write, prefix, flag)
	return &Logger{newLogger: l}
}

func (log *Logger) clone() *Logger {
	nl := *log
	return &nl
}

// 设置日志公共字段
func (log *Logger) WithFields(f Fields) *Logger {
	ll := log.clone()
	if ll.fields != nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

// 设置日志上下文属性
func (log *Logger) WithContext(ctx context.Context) *Logger {
	ll := log.clone()
	ll.ctx = ctx
	return ll
}

// 设置当前某一层调用栈的信息
func (log *Logger) WithCaller(skip int) *Logger {
	ll := log.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return ll
}

// 设置当前的整个调用栈信息
func (log *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := log.clone()
	ll.callers = callers
	return ll
}

func (log *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(log.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	return data
}
