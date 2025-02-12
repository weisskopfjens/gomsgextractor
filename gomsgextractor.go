// Copyright 2023 Jens Weißkopf. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gomsgextractor

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf16"

	"github.com/richardlehane/mscfb"
)

// decodeUTF16String decode UTF-16LE-Bytes into a go string
func DecodeUTF16String(b []byte) string {
	if len(b)%2 != 0 {
		b = b[:len(b)-1] // Bei ungerader Länge das letzte Byte abschneiden
	}
	u16s := make([]uint16, len(b)/2)
	for i := 0; i < len(u16s); i++ {
		u16s[i] = binary.LittleEndian.Uint16(b[i*2:])
	}
	s := string(utf16.Decode(u16s))
	return strings.TrimRight(s, "\x00")
}

// ExtractAttachments extract all attachments from a .msg file
// und speichert sie im angegebenen Ausgabeverzeichnis.
func ExtractAttachments(msgFile string, outDir string) error {
	file, _ := os.Open(msgFile)
	defer file.Close()
	doc, err := mscfb.New(file)
	if err != nil {
		return err
	}
	attachments := make(map[string]map[string]*mscfb.File)
	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {
		if len(entry.Path) == 0 {
			continue
		}
		if attachments[entry.Path[0]] == nil {
			attachments[entry.Path[0]] = make(map[string]*mscfb.File)
		}
		attachments[entry.Path[0]][entry.Name] = entry
	}

	for attachStorage, streams := range attachments {
		if strings.HasPrefix(attachStorage, "__attach_version1.0_") {
			filename := ""
			fmt.Printf("Storage %s\n", attachStorage)
			// retrieve filename of the attachment
			if entry, ok := streams["__substg1.0_3704001E"]; ok {
				buf := make([]byte, 512)
				i, _ := entry.Read(buf)
				if i > 0 {
					out := DecodeUTF16String(buf[:i])
					fmt.Printf("%s = [%s]\n", entry.Name, out)
					filename = out
				}
				fmt.Println(entry.Name, i)
			}
			// retrieve filename of the attachment, method 2
			if entry, ok := streams["__substg1.0_3704001F"]; ok {
				buf := make([]byte, 512)
				i, _ := entry.Read(buf)
				if i > 0 {
					out := DecodeUTF16String(buf[:i])
					fmt.Printf("%s = [%s]\n", entry.Name, out)
					filename = out
				}
			}
			fmt.Println(filename)
			// content stream
			if entry, ok := streams["__substg1.0_37010102"]; ok {
				// open output file
				outPath := filepath.Join(outDir, filename)
				fo, err := os.Create(outPath)
				if err != nil {
					panic(err)
				}
				// close fo on exit and check for its returned error
				defer func() {
					if err := fo.Close(); err != nil {
						panic(err)
					}
				}()

				buf := make([]byte, 512)

				for {
					// read a chunk
					n, err := entry.Read(buf)
					if err != nil && err != io.EOF {
						panic(err)
					}
					if n == 0 {
						break
					}

					// write a chunk
					if _, err := fo.Write(buf[:n]); err != nil {
						panic(err)
					}
				}

			}
		}
	}

	return nil
}
