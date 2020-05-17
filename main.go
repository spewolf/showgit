package main

import (
	"flag"
)

// look for repos and add to the list
func scan(path string) {
	print("scan")
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
