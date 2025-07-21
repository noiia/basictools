// Copyright (c) 2024–2025 Edwin Lecomte
// This file is licensed under the MIT License.
// See the LICENSE file in the root of this repository.

package logger_test

import (
	"os"
	"testing"

	"github.com/noiia/basictool/errornow"
	"github.com/noiia/basictool/logger"

	"github.com/noiia/basictool/editstring"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	writeLogers, err := logger.Logger(false)
	if err != nil {
		errornow.KillComment(t, err)
	}

	writeLogers.Error("verbose logger test")

	body, err := os.ReadFile("logs/logs.log")
	if err != nil {
		errornow.KillComment(t, err)
	}

	if editstring.Clean(string(body)) != "" {
		t.Log("non verbose logger test : success")
	}
}
