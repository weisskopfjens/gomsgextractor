// Copyright 2023 Jens Weißkopf. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// Forked, changed and extend from https://github.com/JacThomp/docstract
// Special thx to JacThomp

package docstract

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

	startdata := ((len(*(d.FileName)) + 2) * 2)
	m := []byte(d.spreadString(d.Mime))
	e := []byte(d.spreadString(filepath.Ext(*d.FileName)))
	f := []byte(d.spreadString(*d.FileName))
	t := d.Bytes

	fmt.Println(string(m))
	fmt.Println(string(e))
	fmt.Println(string(f))

	p1 := bytes.LastIndex(t, m)
	if p1 == -1 {
		return errors.New("Mime not found at end of data.")
	}
	p2 := p1 - len(f) - len(e)
	enddata := bytes.LastIndex(t[:p2], e)
	if enddata == -1 {
		return errors.New("Ext not found at end of data.")
	}

	_, err := os.Stat(path + *(d.FileName))
	if os.IsNotExist(err) {
		return os.WriteFile(path+*(d.FileName), d.Bytes[startdata:enddata], 0644)
	}

	for i := 1; true; i++ {
		istr := strconv.Itoa(i)
		_, err := os.Stat(path + istr + "_" + *d.FileName)
		if os.IsNotExist(err) {
			return os.WriteFile(path+istr+"_"+*d.FileName, d.Bytes[startdata:], 0644)
		}
	}
	return fmt.Errorf("impossible error")
}

func (d *DocStract) spreadString(str string) string {
	var o []rune
	for _, v := range str {
		o = append(o, v)
		o = append(o, 0)
	}
	return string(o)
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

// exchangeInvisibleChars make from a invisible char a visible _
func (d *DocStract) exchangeInvisibleChars(str string) string {
	clean := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return ' '
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
