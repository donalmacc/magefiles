//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	// mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	fmt.Println("Building...")
	return nil
}

// Cleans the stuff
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}
