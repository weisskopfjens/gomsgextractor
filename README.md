# goMSGExtractor
With this command line tool you can extract attachments from microsoft outlook .msg files.

I completely rewrite the code. Now it uses the "github.com/richardlehane/mscfb" package.

The old version use parts of github.com/JacThomp/docstract. Now it was removed from this project.

## Dependencies
github.com/richardlehane/mscfb

## Build
For linux binaries you can use the ```./build.sh``` script.

Alternatively:
In the cmd directory type ```go build -o gomsgextractor```

## Usage
```
Usage: gomsgextractor <msg-file> <output-dir>
```

## Todo
- Write unit tests
- Extract from multiple files