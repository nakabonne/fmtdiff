# fmtdiff

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nakabonne/fmtdiff)

A `goimports` client as well as a parser that parses the diff between an original file and one formatted by it.  

goimports not only fixes imports, but also formats your code in the same style as gofmt, so `fmtdiff` means `importsdiff` substantially.

## Installation

```
go get github.com/nakabonne/fmtdiff
```

## Usage Example

```go
package main	

import "github.com/nakabonne/fmtdiff"

func main() {	
	fileDiff, _ := fmtdiff.Run("/path/to/foo.go", &fmtdiff.Options{
		LocalPrefixes:  []string{"github.com/myOrg/myRepo"},
                IgnoreComments: true,
                FormatOnly:     true,
	})	
}
```

## Thanks

Thanks to [sourcegraph/go-diff](https://github.com/sourcegraph/go-diff) for cool diff parser.
