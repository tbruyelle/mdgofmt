package main

import (
	"fmt"
	"io"
	"os"

	"github.com/tbruyelle/mdgofmt"
)

func main() {
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bz, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	bz, err = mdgofmt.Format(bz)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(bz))
}
