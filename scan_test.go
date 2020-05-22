package main

import (
	"os"
	"testing"
)

func TestRecursiveScan(t *testing.T) {
	wd, _ := os.Getwd()
	repos := recursiveScanFolder(wd)

	if len(repos) != 1 {
		t.Errorf("Recursive search failed in project folder\n")
	}
}

func TestRecursiveScanCapture(t *testing.T) {
	wd, _ := os.Getwd()
	repos := recursiveScanFolder(wd)

	if repos[0] != wd {
		t.Errorf("Recursive found %v and expected %v\n", repos[0], wd)
	}
}

func TestSliceContains(t *testing.T) {
	x := []string{"hi", "i", "am", "test"}
	if !sliceContains(x, "i") {
		t.Errorf("Expected true recieved false")
	}
}

func TestSliceContains_False(t *testing.T) {
	x := []string{"hi", "i", "am", "test"}
	if sliceContains(x, "parrot") {
		t.Errorf("Expected false recieved true")
	}
}

func TestJoinSlices(t *testing.T) {
	a := []string{"hello", "i", "am"}
	b := []string{"i", "am", "a program"}

	c := joinSlices(a, b)

	if len(c) != 4 {
		t.Errorf("Expected 4 elements")
	}
}
