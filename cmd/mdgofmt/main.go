package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tbruyelle/mdgofmt"
)

var write = flag.Bool("w", false, "write result to (source) file instead of stdout")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: mdgofmt [flags] path ...\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}
	for _, file := range flag.Args() {
		bz, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "can't read file %s: %v\n", file, err)
			os.Exit(1)
		}
		bz, err = mdgofmt.Format(bz)
		if err != nil {
			fmt.Fprintf(os.Stderr, "can't format file %s: %v\n", file, err)
			os.Exit(1)
		}
		if *write {
			os.WriteFile(file, bz, 0o644)
		} else {
			fmt.Print(string(bz))
		}
	}
}
