package main

var (
	GIT_COMMIT string
	GIT_TAG    string
	BUILD_DATE string
)

func GetVersion() string {
	return GIT_TAG + " " + GIT_COMMIT + " " + BUILD_DATE
}
