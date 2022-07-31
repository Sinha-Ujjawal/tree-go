package main

import (
	"flag"
	"os"
	"tree/tree"
)

type flags struct {
	dirPath string
}

func parseFlags() flags {
	dirPath := flag.String("dirPath", ".", "Path of the directory from where to start recursing")
	flag.Parse()
	return flags{dirPath: *dirPath}
}

func main() {
	flgs := parseFlags()
	t, err := tree.New(&flgs.dirPath)
	if err != nil {
		panic(err)
	}
	err = t.PrintTree(os.Stdout)
	if err != nil {
		panic(err)
	}
}
