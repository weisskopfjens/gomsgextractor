// Copyright 2025 Jens Weißkopf. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/weisskopfjens/gomsgextractor"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <msg-file> <output-dir>\n", os.Args[0])
		os.Exit(1)
	}
	msgFile := os.Args[1]
	outDir := os.Args[2]

	err := gomsgextractor.ExtractAttachments(msgFile, outDir)
	if err != nil {
		fmt.Printf("Fehler: %v\n", err)
		os.Exit(1)
	}
}
