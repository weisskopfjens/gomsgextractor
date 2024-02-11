// Copyright 2023 Jens Weißkopf. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/weisskopfjens/gomsgextractor/docstract"
)

func main() {
	var filename, outputdir string
	var err error
	flag.StringVar(&filename, "file", "", "(*) A .msg file")
	flag.StringVar(&outputdir, "out", "", "Specify a custom output directory.")
	flag.Parse()
	if filename == "" {
		fmt.Println("goMSGExtractor by (c)Jens Weißkopf (github.com/weisskopfjens/gomsgextractor)")
		fmt.Println("(*) are required parameter.")
		flag.Usage()
		os.Exit(0)
	}
	file, _ := os.ReadFile(filename)

	files, count, err := docstract.Extract(file)

	if err != nil {
		panic(err)
	}

	fmt.Println("Found ", count, " attached files.")

	for _, document := range *files {
		if document.FileName == nil {
			continue
		}
		if err := document.SaveFile(outputdir); err != nil {
			log.Fatalln(err)
		} else {
			fmt.Printf("Saved file %s", *document.FileName)
		}

	}

}
