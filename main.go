package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := []string{"."}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		fmt.Println(arg)
		err := tree(arg)
		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}
}

func tree(root string) error {
	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.Name()[0] == '.' {
			return filepath.SkipDir
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return fmt.Errorf("could not rel(%s, %s): %v", root, path, err)
		}
		depth := len(strings.Split(rel, string(filepath.Separator)))
		fmt.Printf("%s%s\n", strings.Repeat("  ", depth), fi.Name())
		return nil
	})
	return err
}
