# fmtdiff

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nakabonne/fmtdiff)

A `goimports` client as well as a parser that parses the diff between an original file and a formatted one.  

goimports not only fixes imports, but also formats your code in the same style as gofmt, so `fmtdiff` means `importsdiff` substantially.

## Installation

```
go get github.com/nakabonne/fmtdiff
```

## Usage Example

```go
package main	

import "github.com/nakabonne/fmtreporter"	

func main() {	
	fileDiff, _ := fmtreporter.Run("/path/to/foo.go", &fmtreporter.Options{	
		LocalPrefix: "github.com/orgA/repoB",	
		Fragment:   true,	
		TabWidth:   8,	
		TabIndent:  true,	
		FormatOnly: true,
	})	
}
```
