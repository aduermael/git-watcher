package main

import (
	"fmt"
	"log"
	"os"
)

func fail(e error) {
	log.Fatalf("error: %v", e)
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
