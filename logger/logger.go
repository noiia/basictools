// Copyright (c) 2024–2025 Edwin Lecomte
// This file is licensed under the MIT License.
// See the LICENSE file in the root of this repository.

package logger

import (
	"io"
	"log/slog"
	"os"
)

const logPermissions = 0o644

var GlobalLogger *slog.Logger

func InitLogger(verbose bool) error {
	var loggerErr error
	GlobalLogger, loggerErr = Logger(verbose)
	if loggerErr != nil {
		return loggerErr
	}
	return nil
}

// Return a new writer used as logger, create a folder and a logs.log file if it doesn't exist.
//
// If true argument is given, it creates a multiwritter that writes in log file and writes in the terminal.
//
// Else it creates a lonely writer and writes in log file.
func Logger(isVerbose bool) (*slog.Logger, error) {
	var Logfile *os.File

	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		if errMkdr := os.Mkdir("logs", os.ModePerm); errMkdr != nil {
			return nil, errMkdr
		}
	}

	if Logfile, err = os.OpenFile("logs/logs.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, logPermissions); err != nil {
		return nil, err
	}

	var multiWriter io.Writer

	if isVerbose {
		multiWriter = io.MultiWriter(Logfile, os.Stdout)
	} else {
		multiWriter = Logfile
	}

	return slog.New(slog.NewJSONHandler(multiWriter, nil)), nil
}
