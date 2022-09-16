package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tbruyelle/mdgofmt"
)

var write = flag.Bool("w", false, "write result to (source) file instead of stdout")

func main() {
	flag.Parse()
	for _, file := range flag.Args() {
		bz, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		bz, err = mdgofmt.Format(bz)
		if err != nil {
			panic(err)
		}
		if *write {
			os.WriteFile(file, bz, 0o644)
		} else {
			fmt.Print(string(bz))
		}
	}
}
