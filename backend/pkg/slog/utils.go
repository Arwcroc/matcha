package slog

import "os"

func LogError(err error) {
	if err != nil {
		Error(err)
	}
}

func LogErrorExit(err error) {
	if err != nil {
		Error(err)
		os.Exit(2)
	}
}
