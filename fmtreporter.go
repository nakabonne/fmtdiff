package fmtreporter

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/tools/imports"

	"github.com/nakabonne/fmtreporter/diff"
)

type FileDiff struct {
	Name   string
	Before []byte
	After  []byte

	Hunks []*Hunk
}

func (f *FileDiff) NoIssue() bool {
	return bytes.Equal(f.Before, f.After)
}

type Hunk struct {
	// starting line number in original file
	OrigStartLine int
	// number of lines the hunk applies to in the original file
	OrigLines int
	// if > 0, then the original file had a 'No newline at end of file' mark at this offset
	OrigNoNewlineAt int
	// starting line number in new file
	NewStartLine int
	// number of lines the hunk applies to in the new file
	NewLines int
	// optional section heading
	Section string
	// 0-indexed line offset in unified file diff (including section headers); this is
	// only set when Hunks are read from entire file diff (i.e., when ReadAllHunks is
	// called) This accounts for hunk headers, too, so the StartPosition of the first
	// hunk will be 1.
	StartPosition int
	// hunk body (lines prefixed with '-', '+', or ' ')
	Body []byte
}

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

var defaultOption = &Options{
	Fragment:   true,
	TabWidth:   8,
	TabIndent:  true,
	FormatOnly: true,
}

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

	if fileDiff.NoIssue() {
		return fileDiff, nil
	}

	d, err := diff.Diff(src, res, filename)
	if err != nil {
		return nil, err
	}
	fileDiff.Hunks = make([]*Hunk, 0, len(d.Hunks))
	for _, h := range d.Hunks {
		fileDiff.Hunks = append(fileDiff.Hunks, &Hunk{
			OrigStartLine:   int(h.OrigStartLine),
			OrigLines:       int(h.OrigLines),
			OrigNoNewlineAt: int(h.OrigNoNewlineAt),
			NewStartLine:    int(h.NewStartLine),
			NewLines:        int(h.NewLines),
			Section:         h.Section,
			StartPosition:   int(h.StartPosition),
			Body:            h.Body,
		})
	}

	return fileDiff, nil
}
