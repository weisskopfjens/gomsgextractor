// Copyright 2023 Jens Weißkopf. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package docmime

import (
	"encoding/json"
)

// Class to handle mimetypes and corresponding file extension
type DocMime struct {
	mimelist []Mime
}

type Mime struct {
	Extension      string `json:"Extension"`
	MIMEType       string `json:"MIME Type"`
	KindOfDocument string `json:"Kind of document"`
}

// Init must first call befor use
func (d *DocMime) Init() error {
	err := json.Unmarshal([]byte(docmimejson), &d.mimelist)
	return err
}

// Return a string array with all known mimetypes
func (d *DocMime) GetAllMimetypes() []string {
	var o []string
	for _, v := range d.mimelist {
		o = append(o, v.MIMEType)
	}
	return o
}

// GetExtByMimetype Return fileextension for corresponding mimetype
// Return an empty string for unknown mimetype
func (d *DocMime) GetExtByMimetype(mimetype string) string {
	for _, v := range d.mimelist {
		if v.MIMEType == mimetype {
			return v.Extension
		}
	}
	return ""
}

// GetMimetypeByExt Return mimetype for given fileexteension
// Return an empty string for unknown fileextension
func (d *DocMime) GetMimetypeByExt(ext string) string {
	for _, v := range d.mimelist {
		if v.Extension == ext {
			return v.MIMEType
		}
	}
	return ""
}

// docmimejson hold a json encoded list of known mimetype
var docmimejson string = `[
    {
        "Extension": ".aac",
        "MIME Type": "audio/aac",
        "Kind of document": "AAC audio"
    },
    {
        "Extension": ".abw",
        "MIME Type": "application/x-abiword",
        "Kind of document": "AbiWord document"
    },
    {
        "Extension": ".apng",
        "MIME Type": "image/apng",
        "Kind of document": "Animated Portable Network Graphics (APNG) image"
    },
    {
        "Extension": ".arc",
        "MIME Type": "application/x-freearc",
        "Kind of document": "Archive document (multiple files embedded)"
    },
    {
        "Extension": ".avif",
        "MIME Type": "image/avif",
        "Kind of document": "AVIF image"
    },
    {
        "Extension": ".avi",
        "MIME Type": "video/x-msvideo",
        "Kind of document": "AVI: Audio Video Interleave"
    },
    {
        "Extension": ".azw",
        "MIME Type": "application/vnd.amazon.ebook",
        "Kind of document": "Amazon Kindle eBook format"
    },
    {
        "Extension": ".bin",
        "MIME Type": "application/octet-stream",
        "Kind of document": "Any kind of binary data"
    },
    {
        "Extension": ".bmp",
        "MIME Type": "image/bmp",
        "Kind of document": "Windows OS/2 Bitmap Graphics"
    },
    {
        "Extension": ".bz",
        "MIME Type": "application/x-bzip",
        "Kind of document": "BZip archive"
    },
    {
        "Extension": ".bz2",
        "MIME Type": "application/x-bzip2",
        "Kind of document": "BZip2 archive"
    },
    {
        "Extension": ".cda",
        "MIME Type": "application/x-cdf",
        "Kind of document": "CD audio"
    },
    {
        "Extension": ".csh",
        "MIME Type": "application/x-csh",
        "Kind of document": "C-Shell script"
    },
    {
        "Extension": ".css",
        "MIME Type": "text/css",
        "Kind of document": "Cascading Style Sheets (CSS)"
    },
    {
        "Extension": ".csv",
        "MIME Type": "text/csv",
        "Kind of document": "Comma-separated values (CSV)"
    },
    {
        "Extension": ".doc",
        "MIME Type": "application/msword",
        "Kind of document": "Microsoft Word"
    },
    {
        "Extension": ".docx",
        "MIME Type": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        "Kind of document": "Microsoft Word (OpenXML)"
    },
    {
        "Extension": ".eot",
        "MIME Type": "application/vnd.ms-fontobject",
        "Kind of document": "MS Embedded OpenType fonts"
    },
    {
        "Extension": ".epub",
        "MIME Type": "application/epub+zip",
        "Kind of document": "Electronic publication (EPUB)"
    },
    {
        "Extension": ".gz",
        "MIME Type": "application/gzip",
        "Kind of document": "GZip Compressed Archive"
    },
    {
        "Extension": ".gif",
        "MIME Type": "image/gif",
        "Kind of document": "Graphics Interchange Format (GIF)"
    },
    {
        "Extension": ".html",
        "MIME Type": "text/html",
        "Kind of document": "HyperText Markup Language (HTML)"
    },
    {
        "Extension": ".ico",
        "MIME Type": "image/vnd.microsoft.icon",
        "Kind of document": "Icon format"
    },
    {
        "Extension": ".ics",
        "MIME Type": "text/calendar",
        "Kind of document": "iCalendar format"
    },
    {
        "Extension": ".jar",
        "MIME Type": "application/java-archive",
        "Kind of document": "Java Archive (JAR)"
    },
    {
        "Extension": ".jpg",
        "MIME Type": "image/jpeg",
        "Kind of document": "JPEG images"
    },
    {
        "Extension": ".js",
        "MIME Type": "text/javascript",
        "Kind of document": "JavaScript"
    },
    {
        "Extension": ".json",
        "MIME Type": "application/json",
        "Kind of document": "JSON format"
    },
    {
        "Extension": ".jsonld",
        "MIME Type": "application/ld+json",
        "Kind of document": "JSON-LD format"
    },
    {
        "Extension": ".mid",
        "MIME Type": "audio/midi",
        "Kind of document": "Musical Instrument Digital Interface (MIDI)"
    },
    {
        "Extension": ".mjs",
        "MIME Type": "text/javascript",
        "Kind of document": "JavaScript module"
    },
    {
        "Extension": ".mp3",
        "MIME Type": "audio/mpeg",
        "Kind of document": "MP3 audio"
    },
    {
        "Extension": ".mp4",
        "MIME Type": "video/mp4",
        "Kind of document": "MP4 video"
    },
    {
        "Extension": ".mpeg",
        "MIME Type": "video/mpeg",
        "Kind of document": "MPEG Video"
    },
    {
        "Extension": ".mpkg",
        "MIME Type": "application/vnd.apple.installer+xml",
        "Kind of document": "Apple Installer Package"
    },
    {
        "Extension": ".odp",
        "MIME Type": "application/vnd.oasis.opendocument.presentation",
        "Kind of document": "OpenDocument presentation document"
    },
    {
        "Extension": ".ods",
        "MIME Type": "application/vnd.oasis.opendocument.spreadsheet",
        "Kind of document": "OpenDocument spreadsheet document"
    },
    {
        "Extension": ".odt",
        "MIME Type": "application/vnd.oasis.opendocument.text",
        "Kind of document": "OpenDocument text document"
    },
    {
        "Extension": ".oga",
        "MIME Type": "audio/ogg",
        "Kind of document": "OGG audio"
    },
    {
        "Extension": ".ogv",
        "MIME Type": "video/ogg",
        "Kind of document": "OGG video"
    },
    {
        "Extension": ".ogx",
        "MIME Type": "application/ogg",
        "Kind of document": "OGG"
    },
    {
        "Extension": ".opus",
        "MIME Type": "audio/opus",
        "Kind of document": "Opus audio"
    },
    {
        "Extension": ".otf",
        "MIME Type": "font/otf",
        "Kind of document": "OpenType font"
    },
    {
        "Extension": ".png",
        "MIME Type": "image/png",
        "Kind of document": "Portable Network Graphics"
    },
    {
        "Extension": ".pdf",
        "MIME Type": "application/pdf",
        "Kind of document": "Adobe Portable Document Format (PDF)"
    },
    {
        "Extension": ".php",
        "MIME Type": "application/x-httpd-php",
        "Kind of document": "Hypertext Preprocessor (Personal Home Page)"
    },
    {
        "Extension": ".ppt",
        "MIME Type": "application/vnd.ms-powerpoint",
        "Kind of document": "Microsoft PowerPoint"
    },
    {
        "Extension": ".pptx",
        "MIME Type": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
        "Kind of document": "Microsoft PowerPoint (OpenXML)"
    },
    {
        "Extension": ".rar",
        "MIME Type": "application/vnd.rar",
        "Kind of document": "RAR archive"
    },
    {
        "Extension": ".rtf",
        "MIME Type": "application/rtf",
        "Kind of document": "Rich Text Format (RTF)"
    },
    {
        "Extension": ".sh",
        "MIME Type": "application/x-sh",
        "Kind of document": "Bourne shell script"
    },
    {
        "Extension": ".svg",
        "MIME Type": "image/svg+xml",
        "Kind of document": "Scalable Vector Graphics (SVG)"
    },
    {
        "Extension": ".tar",
        "MIME Type": "application/x-tar",
        "Kind of document": "Tape Archive (TAR)"
    },
    {
        "Extension": ".tif",
        "MIME Type": "image/tiff",
        "Kind of document": "Tagged Image File Format (TIFF)"
    },
    {
        "Extension": ".ts",
        "MIME Type": "video/mp2t",
        "Kind of document": "MPEG transport stream"
    },
    {
        "Extension": ".ttf",
        "MIME Type": "font/ttf",
        "Kind of document": "TrueType Font"
    },
    {
        "Extension": ".txt",
        "MIME Type": "text/plain",
        "Kind of document": "Text, (generally ASCII or ISO 8859-n)"
    },
    {
        "Extension": ".vsd",
        "MIME Type": "application/vnd.visio",
        "Kind of document": "Microsoft Visio"
    },
    {
        "Extension": ".wav",
        "MIME Type": "audio/wav",
        "Kind of document": "Waveform Audio Format"
    },
    {
        "Extension": ".weba",
        "MIME Type": "audio/webm",
        "Kind of document": "WEBM audio"
    },
    {
        "Extension": ".webm",
        "MIME Type": "video/webm",
        "Kind of document": "WEBM video"
    },
    {
        "Extension": ".webp",
        "MIME Type": "image/webp",
        "Kind of document": "WEBP image"
    },
    {
        "Extension": ".woff",
        "MIME Type": "font/woff",
        "Kind of document": "Web Open Font Format (WOFF)"
    },
    {
        "Extension": ".woff2",
        "MIME Type": "font/woff2",
        "Kind of document": "Web Open Font Format (WOFF)"
    },
    {
        "Extension": ".xhtml",
        "MIME Type": "application/xhtml+xml",
        "Kind of document": "XHTML"
    },
    {
        "Extension": ".xls",
        "MIME Type": "application/vnd.ms-excel",
        "Kind of document": "Microsoft Excel"
    },
    {
        "Extension": ".xlsx",
        "MIME Type": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
        "Kind of document": "Microsoft Excel (OpenXML)"
    },
    {
        "Extension": ".xml",
        "MIME Type": "application/xml",
        "Kind of document": "XML"
    },
    {
        "Extension": ".xul",
        "MIME Type": "application/vnd.mozilla.xul+xml",
        "Kind of document": "XUL"
    },
    {
        "Extension": ".zip",
        "MIME Type": "application/zip",
        "Kind of document": "ZIP archive"
    },
    {
        "Extension": ".3gp",
        "MIME Type": "video/3gpp; audio/3gpp",
        "Kind of document": "3GPP audio/video container"
    },
    {
        "Extension": ".3g2",
        "MIME Type": "video/3gpp2; audio/3gpp2",
        "Kind of document": "3GPP2 audio/video container"
    },
    {
        "Extension": ".7z",
        "MIME Type": "application/x-7z-compressed",
        "Kind of document": "7-zip archive"
    }
]
`
