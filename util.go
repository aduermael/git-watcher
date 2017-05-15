package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func fail(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

func rmDir(path string) {
	debug("rm", path)
	err := os.RemoveAll(path)
	if err != nil {
		fail(err)
	}
}

func debug(args ...interface{}) {
	fmt.Println(args...)
}

func filePathMatchPattern(pattern, path string) bool {
	// if pattern is of this form: *.png (with no '/')
	// then it means we want to check the file name and not complete path
	if !filepath.IsAbs(pattern) && filepath.Base(pattern) == pattern {
		debug("looking for file pattern")
		path = filepath.Base(path)
	}

	matched, err := filepath.Match(pattern, path)
	if err != nil {
		// TODO: proper error handling
		return false
	}

	return matched
}
