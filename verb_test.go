// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package main

import "testing"

func TestGuessVerb(t *testing.T) {
	verbs := map[string]string{"werden": "diventare"}
	got := GuessVerb("diventare", "werden", verbs) 
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
