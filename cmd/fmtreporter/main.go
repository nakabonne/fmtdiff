package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/nakabonne/fmtreporter"
)

var (
	flagSet     = flag.NewFlagSet("fmtreporter", flag.ContinueOnError)
	localPrefix = flagSet.String("local-prefix", "", "put imports beginning with this string after 3rd-party packages; comma-separated list.")
	formatOnly  = flagSet.Bool("format-only", false, "if true, don't fix imports and only format. In this mode, goimports is effectively gofmt, with the addition that imports are grouped into sections.")
)

func main() {
	flagSet.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: fmtreporter [flags] [files ...]")
		flagSet.PrintDefaults()
	}
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	fileDiffs := []*fmtreporter.FileDiff{}
	for _, path := range flagSet.Args() {
		fs, err := fmtreporter.Run(path, &fmtreporter.Options{
			LocalPrefix: *localPrefix,
			// TODO: Make these configuable as well.
			Fragment:   true,
			TabWidth:   8,
			TabIndent:  true,
			FormatOnly: *formatOnly,
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
		fileDiffs = append(fileDiffs, fs)
	}
	for _, f := range fileDiffs {
		pp.Println(f)
	}
}
