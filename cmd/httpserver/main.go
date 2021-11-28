package main

import (
	"flag"
	"fmt"
	"os"

	"cnjob/internal/command"
	"cnjob/internal/logger"
)

var (
	Listener    string
	Version     string
	LogLevel    string
	LogPath     string
	LogCaller   bool
	showVersion bool
)

func init() {
	flag.StringVar(&Listener, "listen", ":80", "set listen")
	flag.StringVar(&LogLevel, "log_level", "warn", fmt.Sprintf("support log level: \n\t%v", logger.GetSupportLogLevelToString()))
	flag.StringVar(&LogPath, "log_path", "", "set log path")
	flag.BoolVar(&LogCaller, "log_caller", true, "print log caller")
	flag.BoolVar(&showVersion, "V", false, "Show version Details")
	flag.Parse()

	if showVersion {
		fmt.Println(GetVersion())
		os.Exit(0)
	}

	Version = os.Getenv("VERSION")
	if Version == "" {
		Version = GetVersion()
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
