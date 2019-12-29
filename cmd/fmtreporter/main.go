package main

import (
	"flag"
	"fmt"
	"os"

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

	issues := []*fmtreporter.Issue{}
	for _, path := range flagSet.Args() {
		is, err := fmtreporter.Run(path, &fmtreporter.Options{
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
		issues = append(issues, is...)
	}
	for _, i := range issues {
		fmt.Println(i.String())
	}
}
