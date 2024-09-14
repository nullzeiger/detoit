// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package main

import (
	"fmt"
	"os"
)

func GuessVerb(g string, k string, verbs map[string]string) bool {
	return verbs[k] == g
}

func PrintVerb() {
	verbs := map[string]string{
		"werden":  "diventare",
		"haben":   "avere",
		"sein":    "essere",
		"können":  "potere",
		"müssen":  "dovere",
		"sollen":  "dovere",
		"sagen":   "dire",
		"geben":   "dare",
		"kommen":  "venire",
		"wollen":  "volere",
		"machen":  "fare",
		"gehen":   "andare",
		"heißen":  "chiamare",
		"wisen":   "sapere",
		"sehen":   "vedere",
		"finden":  "trovare",
		"bleiben": "rimanere",
		"mögen":   "piacere",
		"fahren":  "andare",
	}

	right := 0

	for k, v := range verbs {
		var guess string

		fmt.Printf("Traduci verbo %s in Italiano (q per uscire): ", k)
		fmt.Scanln(&guess)

		if guess == "q" {
			os.Exit(0)
		}

		if GuessVerb(guess, k, verbs) {
			fmt.Printf("Giusto %s %s\n", k, v)
			right++
		} else {
			fmt.Printf("Sbagliato %s %s\n", k, v)
		}
	}

	fmt.Printf("Risposte giuste: %d su %d verbi", right, len(verbs))

}
