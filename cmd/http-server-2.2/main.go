package main

import (
	"cnjob/internal/command"
	"cnjob/internal/logger"
	"flag"
	"fmt"
	"os"
)

var (
	Listener  string
	Version   string
	LogLevel  string
	LogPath   string
	LogCaller bool
)

func init() {
	flag.StringVar(&Listener, "listen", ":80", "set listen")
	flag.StringVar(&Version, "version", "", "set version")
	flag.StringVar(&LogLevel, "log_level", "warn", fmt.Sprintf("support log level: \n\t%v", logger.GetSupportLogLevelToString()))
	flag.StringVar(&LogPath, "log_path", "", "set log path")
	flag.BoolVar(&LogCaller, "log_caller", true, "print log caller")
	flag.Parse()

	if Version == "" {
		Version = os.Getenv("VERSION")
	}

	Level, err := logger.ParseLevel(LogLevel)
	if err != nil {
		panic(err)
	}

	logger.Init(
		logger.WithLevel(Level),
		logger.WithPath(LogPath),
		logger.WithReportCaller(LogCaller),
	)
}

func main() {
	command.Run(NewApp())
}
