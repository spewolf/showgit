package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func trimSuffixes(str string, suffixes []string) string {
	if len(suffixes) > 0 {
		str = strings.TrimSuffix(str, suffixes[0])
		str = trimSuffixes(str, suffixes[1:])
	}
	return str
}

func recursiveScanFolder(path string) []string {
	repos := make([]string, 0)

	// walks file structure recursivly and adds directories with .git folders
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			// filter out bulky folders without repos inside
			if info.Name() == "node_modules" || info.Name() == "vendor" {
				fmt.Printf("Skipping folder located at: %v\n", path)
				return filepath.SkipDir
			} else if info.Name() == ".git" {
				// add folders containing .git directories
				repos = append(repos, trimSuffixes(path, []string{"\\.git", "/.git"}))
			}
		}

		return nil
	})

	return repos
}
