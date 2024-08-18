// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package main

import (
	"fmt"
	"os"
)

func main() {
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

	for key, value := range verbs {
		var guess string

		fmt.Printf("Traduci verbo %s in Italiano (q per uscire): ", key)
		fmt.Scanln(&guess)

		if guess == "q" {
			os.Exit(0)
		}

		if verbs[key] == guess {
			fmt.Printf("Giusto %s %s\n", key, value)
		} else {
			fmt.Printf("Sbagliato %s %s\n", key, value)
		}
	}
}
