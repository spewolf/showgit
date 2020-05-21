package main

import (
	"bufio"
	"fmt"
	"io"
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

func combineSliceWithFile(repos []string, path string) {
	// existingRepos := parseFileLinesToSlice(path)
	// join slices
	// write to file
}

func parseFileLinesToSlice(path string) []string {
	file := openFile(path)
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		panic(err)
	}

	return lines
}

func openFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(path)
			if err != nil {
				panic(err)
			} else {
				panic(err)
			}
		}
	}
	return f
}
