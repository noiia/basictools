// Copyright (c) 2024–2025 Edwin Lecomte
// This file is licensed under the MIT License.
// See the LICENSE file in the root of this repository.

package errornow_test

import (
	"testing"

	"github.com/noiia/basictools/errornow"
)

type testKillComment struct {
	wellLogged bool
	wellFailed bool
}

func (k *testKillComment) Log(args ...any) {
	k.wellLogged = true
}

func (k *testKillComment) FailNow() {
	k.wellFailed = true
}

func TestErrorNow(t *testing.T) {
	t.Parallel()

	newKillComment := testKillComment{
		wellLogged: false,
		wellFailed: false,
	}

	errornow.KillComment(&newKillComment, "test error comments")

	if newKillComment.wellLogged && newKillComment.wellFailed {
		t.Log("kill comment test : success")
	} else {
		t.Error("test kill comment test : failure")
		t.FailNow()
	}
}
