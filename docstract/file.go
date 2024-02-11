// Copyright 2023 Jens Weißkopf. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// Forked, changed and extend from https://github.com/JacThomp/docstract
// Special thx to JacThomp

package docstract

import (
	"errors"
	"os"
	"strings"
	"unicode"

	"github.com/weisskopfjens/gomsgextractor/docmime"
)

// DocStract stores the binary data for extracted files, as well as the type and filename metadata
type DocStract struct {
	Mime     string
	FileName *string
	Bytes    []byte
}

// SaveFile saves the file to the path, does not check if it's an unkown filetype only if it has a name
func (d *DocStract) SaveFile(path string) error {
	if len(path) > 0 && path[len(path)-1] != '/' {
		path += "/"
	}

	if d.FileName == nil {
		return errors.New("document does not have a filename cannot save")
	}

	return os.WriteFile(path+*(d.FileName), d.Bytes[((len(*(d.FileName))+2)*2):], 0644)
}

func (d *DocStract) removeInvisibleChars(str string) string {
	clean := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)
	return clean
}

// Helper for getName DRY ext must be fileextension with point .
// mime for example application/pdf
func (d *DocStract) getNameHelper(block string, mime string, ext string) string {
	if strings.Contains(block, mime) {
		d.Mime = mime
		end := strings.Index(block, mime)
		start := 0
		//fmt.Println("end:", end, "ext:", len(ext))
		for i := end - len(ext) - 1; i > 0; i-- {
			if strings.Contains(block[i:i+len(ext)], ext) {
				start = i + len(ext)
				return block[start:end]
			}
		}
	}
	return ""
}

// sets name to nil if cannot dertermine name and type to unkown
func (d *DocStract) getName() {
	blocks := strings.Split(string(d.Bytes), "\n")
	nameBlock := blocks[len(blocks)-1]
	nameBlock = d.removeInvisibleChars(nameBlock)
	mime := docmime.DocMime{}

	mime.Init()
	d.FileName = nil
	d.Mime = "unknown"
	// Get string array of all mime tyoes
	allmimes := mime.GetAllMimetypes()
	for _, v := range allmimes {
		// Get fileextension by mime type
		e := mime.GetExtByMimetype(v)
		// Look in the nameBlock and analyse if the mimetype is in it
		s := d.getNameHelper(nameBlock, v, e)
		if s == "" {
			// not the right mime type. go on.
			continue
		}
		// mimetype found
		d.FileName = &s
		d.Mime = v
		break
	}
}
