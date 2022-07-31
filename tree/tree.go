package tree

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Tree struct {
	rootPath *string
}

func New(rootPath *string) (*Tree, error) {
	path, err := os.Stat(*rootPath)
	if err != nil {
		return nil, err
	}
	if !path.IsDir() {
		return nil, errors.New("Make sure that the path is a directory")
	}
	return &Tree{rootPath}, nil
}

func (tree *Tree) PrintTree(w io.Writer) error {
	return printTreeRecurse(w, 0, tree.rootPath)
}

func printTreeRecurse(w io.Writer, indents uint, dirPath *string) error {
	entries, err := ioutil.ReadDir(*dirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fileName := entry.Name()
		prefix := strings.Repeat("  ", int(indents))
		fmt.Fprintf(w, "%s|-", prefix)
		fmt.Fprintln(w, fileName)
		if entry.IsDir() {
			newDirPath := path.Join(*dirPath, entry.Name())
			err := printTreeRecurse(w, indents+1, &newDirPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
