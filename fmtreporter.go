package fmtreporter

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/tools/imports"
)

var (
	options = &imports.Options{
		TabWidth:   8,
		TabIndent:  true,
		Comments:   true,
		Fragment:   true,
		FormatOnly: true,
	}
)

func Run(filename, localPrefix string) ([]byte, error) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	imports.LocalPrefix = localPrefix
	res, err := imports.Process(filename, src, options)
	if err != nil {
		return nil, err
	}

	if bytes.Equal(src, res) {
		fmt.Println("the two files are the same length and contain the same bytes")
		return nil, nil
	}

	// formatting has changed
	/*data, err := diff(src, res, filename)
	if err != nil {
		return nil, fmt.Errorf("error computing diff: %s", err)
	}

	return data, nil*/
	return res, nil
}
