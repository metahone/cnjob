package main

import (
	"cnjob/internal/command"
	"cnjob/internal/logger"
	"flag"
	"fmt"
	"os"
)

var (
	Listener string
	Version  string
	LogLevel string
	LogPath  string
)

func init() {
	flag.StringVar(&Listener, "listen", ":80", "set listen")
	flag.StringVar(&Version, "version", "", "set version")
	flag.StringVar(&LogLevel, "log_level", "warn", fmt.Sprintf("support log level: \n\t%v", logger.GetSupportLogLevelToString()))
	flag.StringVar(&LogPath, "log_path", "", "set log path")
	flag.Parse()

	if Version == "" {
		Version = os.Getenv("VERSION")
	}

	Level, err := logger.ParseLevel(LogLevel)
	if err != nil {
		panic(err)
	}
	logger.Init(logger.WithLevel(Level), logger.WithPath(LogPath))
}

func main() {
	command.Run(NewApp())
}
