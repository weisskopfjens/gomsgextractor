// Copyright 2023 Jens Weißkopf. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// Forked, changed and extend from https://github.com/JacThomp/docstract
// Special thx to JacThomp

package docstract

import (
	"bytes"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/richardlehane/mscfb"
)

// Extract takes a .msg files binary data and returns an array of attachments and a count of how many files were extracted
func Extract(data []byte) (*[]*DocStract, int, error) {

	reader := bytes.NewReader(data)

	doc, err := mscfb.New(reader)

	if err != nil {
		return nil, 0, errors.Wrap(err, "creating reader")
	}

	files := []*DocStract{}

	attachment := false
	file := 0
	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {

		if strings.Contains(entry.Name, "attach") {
			files = append(files, &DocStract{})
			attachment = true
			continue
		}
		if attachment && strings.Contains(entry.Name, "properties") {
			attachment = false
			file++
			continue
		}

		if attachment {
			buf := make([]byte, entry.Size)
			i, _ := entry.Read(buf)
			if i > 0 {
				files[file].Bytes = append(files[file].Bytes, buf[:i]...)
				//temp := strconv.Itoa(i) + ".tmp"
				//files[file].FileName = &temp
				//files[file].Type = 0
				//files[file].getName()
			}
		}
	}

	{ //Determine FileType and FileName
		wait := sync.WaitGroup{}
		for _, doc := range files {
			wait.Add(1)
			go func(d *DocStract) {
				d.getName()
				wait.Done()
			}(doc)
		}
		wait.Wait()
	}

	return &files, len(files), nil
}
