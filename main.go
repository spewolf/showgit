package main

import (
	"flag"
	"fmt"
)

// look for repos and add to the list
func scan(path string) {
	fmt.Printf("Found folders:\n\n")
	//repositories := recursiveScanFolder(path)
	// get dot path
	// add new slices
	// print status
}

// print neat contribution graph
func stats(email string) {
	print("stats")
}

func main() {
	var folder, email string
	flag.StringVar(&folder, "add", "", "add a folder to git repositories that are scanned")
	flag.StringVar(&email, "email", "spwolf18@gmail.com", "email that will be searched for")
	flag.Parse()

	if folder != "" {
		scan(folder)
	}

	stats(email)
}
