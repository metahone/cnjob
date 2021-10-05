package logger

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type CustomFormatter struct {
	opt CustomOptions
}

func NewCustomFormatter(opts ...CustomOption) *CustomFormatter {
	l := &CustomFormatter{}
	for _, op := range opts {
		if op == nil {
			continue
		}
		op(&l.opt)
	}
	l.opt.withDefault()
	return l
}

type CustomOptions struct {
	format string
	caller bool
}

type CustomOption func(options *CustomOptions)

func (l *CustomOptions) withDefault() {
	if l.format == "" {
		l.format = "2006-01-02T15:04:05.000Z0700"
	}
}

func WithCustomTimeFormat(format string) CustomOption {
	return func(op *CustomOptions) {
		op.format = format
	}
}
func WithCustomCaller(enable bool) CustomOption {
	return func(op *CustomOptions) {
		op.caller = enable
	}
}

func (s *CustomFormatter) Format(entry *log.Entry) ([]byte, error) {
	//timestamp := time.Now().Local().TimeFormat("0102-150405.000")
	timestamp := time.Now().Local().Format(s.opt.format)

	if s.opt.caller {
		var file string
		var line int
		if entry.Caller != nil {
			file = filepath.Base(entry.Caller.File)
			line = entry.Caller.Line
		}
		//fmt.Println(entry.Data)
		//msg := fmt.Sprintf("%s [%s:%d][GOID:%d][%s] %s\n", timestamp, file, line, getGID(), strings.ToUpper(entry.Level.String()), entry.Message)
		//msg := fmt.Sprintf("%s %s %s:%d [GOID:%d] %s\n", timestamp, strings.ToUpper(entry.Level.String()), file, line, getGID(), entry.Message)
		msg := fmt.Sprintf("%s %s %s:%d %s\n", timestamp, strings.ToUpper(entry.Level.String()), file, line, entry.Message)
		return []byte(msg), nil
	}

	msg := fmt.Sprintf("%s %s %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
