package main

import (
	"fmt"
)

var (
	GitCommit string
	GitTag    string
	BuildDate string
)

func GetVersion() string {
	return fmt.Sprintf("Version info: Tag %v, CommitHash %v, BuildDate %v", GitTag, GitCommit, BuildDate)
}
