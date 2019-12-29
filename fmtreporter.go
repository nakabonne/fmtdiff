package fmtreporter

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/tools/imports"

	"github.com/nakabonne/fmtreporter/diff"
)

// FileDiff represents a diff between an original file and a formatted one.
type FileDiff struct {
	// File name.
	Name string
	// Contents of original file.
	Before []byte
	// Contents of formatted file.
	After []byte
	Hunks []*Hunk
}

// NoDiff checks if the original file and the formatted one
// are the same length and contain the same bytes.
func (f *FileDiff) NoDiff() bool {
	return bytes.Equal(f.Before, f.After)
}

// Hunk represents a series of changes in a file's unified diff.
type Hunk struct {
	// OrigStartLine is the starting line number in the original file.
	OrigStartLine int
	// OrigLines is the number of lines the hunk applies to in the original file.
	OrigLines int
	// NewStartLine is the starting line number in the new file.
	NewStartLine int
	// NewLines is the number of lines the hunk applies to in the new file.
	NewLines int
	Body     []byte
}

// Options makes it possible to fine-tune behavior.
type Options struct {
	// LocalPrefix is a comma-separated string of import path prefixes, which, if
	// set, instructs Process to sort the import paths with the given prefixes
	// into another group after 3rd-party packages.
	LocalPrefix string
	// Accept fragment of a source file (no package statement).
	Fragment bool
	// Use tabs for indent. True is populated if nil provided.
	TabIndent bool
	// 8 is populated if nil provided.
	TabWidth int
	// Disable the insertion and deletion of imports.
	FormatOnly bool
}

var defaultOption = &Options{
	Fragment:  true,
	TabWidth:  8,
	TabIndent: true,
}

// Run runs goimports and parses the diff between an original file and a formatted one.
func Run(filename string, options *Options) (*FileDiff, error) {
	fileDiff := &FileDiff{Name: filename}
	if options == nil {
		options = defaultOption
	}

	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	fileDiff.Before = src

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
	fileDiff.After = res

	if fileDiff.NoDiff() {
		return fileDiff, nil
	}

	d, err := diff.Diff(src, res, filename)
	if err != nil {
		return nil, err
	}
	fileDiff.Hunks = make([]*Hunk, 0, len(d.Hunks))
	for _, h := range d.Hunks {
		fileDiff.Hunks = append(fileDiff.Hunks, &Hunk{
			OrigStartLine: int(h.OrigStartLine),
			OrigLines:     int(h.OrigLines),
			NewStartLine:  int(h.NewStartLine),
			NewLines:      int(h.NewLines),
			Body:          h.Body,
		})
	}

	return fileDiff, nil
}
