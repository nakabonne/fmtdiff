package fmtreporter

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/tools/imports"

	diff "github.com/nakabonne/fmtreporter/diff"
)

var (
	defaultOption = &Options{
		Fragment:   true,
		TabWidth:   8,
		TabIndent:  true,
		FormatOnly: true,
	}
)

type Options struct {
	// LocalPrefix is a comma-separated string of import path prefixes, which, if
	// set, instructs Process to sort the import paths with the given prefixes
	// into another group after 3rd-party packages.
	LocalPrefix string

	Fragment   bool // Accept fragment of a source file (no package statement)
	TabIndent  bool // Use tabs for indent (true if nil *Options provided)
	TabWidth   int  // Tab width (8 if nil *Options provided)
	FormatOnly bool // Disable the insertion and deletion of imports
}

func Run(filename string, options *Options) ([]byte, error) {
	if options == nil {
		options = defaultOption
	}

	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	imports.LocalPrefix = options.LocalPrefix
	res, err := imports.Process(filename, src, &imports.Options{
		Comments:   true,
		Fragment:   options.Fragment,
		TabIndent:  options.TabIndent,
		TabWidth:   options.TabWidth,
		FormatOnly: options.FormatOnly,
	})
	if err != nil {
		return nil, err
	}

	if bytes.Equal(src, res) {
		fmt.Println("the two files are the same length and contain the same bytes")
		return nil, nil
	}

	_, err = diff.Diff(src, res, filename)
	if err != nil {
		return nil, fmt.Errorf("error taking diffs: %s", err)
	}

	return []byte{}, nil
}
