package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/tbruyelle/mdgofmt"
)

var write = flag.Bool("w", false, "write result to (source) file instead of stdout")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: mdgofmt [flags] path\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}
	path := flag.Args()[0]
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read %s: %v\n", path, err)
		os.Exit(1)
	}
	var files []string
	if fi.IsDir() {
		fs.WalkDir(os.DirFS(path), ".", func(p string, d fs.DirEntry, err error) error {
			if !d.Type().IsDir() && filepath.Ext(d.Name()) == ".md" {
				files = append(files, filepath.Join(path, p))
			}
			return nil
		})
	} else {
		files = append(files, path)
	}
	for _, file := range files {
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
