// Copyright (c) 2024–2025 Edwin Lecomte
// This file is licensed under the MIT License.
// See the LICENSE file in the root of this repository.

package errornow

type errorComment interface {
	Log(args ...any)
	FailNow()
}

// allow to create a comment logged and execute the t.FailNow function at the end.
func KillComment(t errorComment, args ...any) {
	t.Log(args...)
	t.FailNow()
}
