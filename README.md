# goMSGExtractor
With this command line tool you can extract attachments from microsoft outlook .msg files.

I completely rewrite the code. Now it uses the "github.com/richardlehane/mscfb" package.

The old version use parts of github.com/JacThomp/docstract. Now it was removed from this project.

## Dependencies
github.com/richardlehane/mscfb

## Usage
```
(*) are required parameter.
Usage of ./gomsgextractor:
  -file string
        (*) A .msg file
  -out string
        Specify a custom output directory.
```

## Todo
- Write unit tests
- Extract from multiple files