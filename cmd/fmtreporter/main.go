package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nakabonne/fmtreporter"
)

var (
	flagSet     = flag.NewFlagSet("fmtreporter", flag.ContinueOnError)
	localPrefix = flagSet.String("local-prefix", "", "local prefix")
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

	for _, path := range flagSet.Args() {
		b, err := fmtreporter.Run(path, &fmtreporter.Options{
			LocalPrefix: *localPrefix,
			// TODO: Make these configuable as well.
			Fragment:   true,
			TabWidth:   8,
			TabIndent:  true,
			FormatOnly: true,
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(b))
	}
}
