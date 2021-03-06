package logger

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"sync"
)

type Options struct {
	TimeFormat   string
	Level        log.Level
	Path         string
	ReportCaller bool
}

type Option func(*Options)

func NewOptions() *Options {
	return &Options{
		Level:        log.Level(999),
		ReportCaller: false,
	}
}

func (o *Options) withDefault() *Options {
	updated := *o
	if updated.Level == log.Level(999) {
		updated.Level = log.WarnLevel
	}
	if updated.TimeFormat == "" {
		//updated.TimeFormat = "2006-01-02T15:04:05.000Z0700"
		updated.TimeFormat = "2006-01-02 15:04:05.000000"
	}
	return &updated
}

func WithLevel(level log.Level) Option {
	return func(args *Options) {
		args.Level = level
	}
}

func WithPath(path string) Option {
	return func(args *Options) {
		args.Path = path
	}
}

func WithReportCaller(report bool) Option {
	return func(args *Options) {
		args.ReportCaller = report
	}
}

func WithTimeFormat(format string) Option {
	return func(args *Options) {
		args.TimeFormat = format
	}
}

func ParseLevel(levelStr string) (log.Level, error) {
	return log.ParseLevel(levelStr)
}

func GetSupportLogLevelToString() string {
	ss := make([]string, 0, len(log.AllLevels))
	for _, v := range log.AllLevels {
		ss = append(ss, v.String())
	}
	return strings.Join(ss, ",")
}

var (
	once sync.Once
)

func Init(opts ...Option) {
	once.Do(func() {
		opt := NewOptions()
		for _, op := range opts {
			if op == nil {
				continue
			}
			op(opt)
		}
		opt = opt.withDefault()

		//log.SetFormatter(&log.JSONFormatter{
		//	TimestampFormat: opt.TimeFormat,
		//})

		customDefaultFormatter := NewCustomFormatter(
			WithCustomTimeFormat(opt.TimeFormat),
			WithCustomCaller(opt.ReportCaller),
		)
		log.SetFormatter(customDefaultFormatter)
		log.SetLevel(opt.Level)
		log.SetReportCaller(opt.ReportCaller)

		var writers []io.Writer

		if len(opt.Path) > 0 {
			logger := &lumberjack.Logger{
				Filename:   opt.Path,
				MaxSize:    128, // megabytes
				MaxAge:     7,   // days
				MaxBackups: 5,
				LocalTime:  true,
				Compress:   true, // disabled by default
			}
			writers = append(writers, logger)
		}

		writers = append(writers, os.Stdout)
		log.SetOutput(io.MultiWriter(writers...))
	})
}
