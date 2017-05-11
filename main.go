package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// WatchConfig is the configuration used to watch Github repositories
type WatchConfig struct {
	Repos map[string]*Repo
}

// Repo represents a watched Github repository
type Repo struct {
	URL      string `yaml:"url"`
	Branches map[string]*Branch
}

type Branch struct {
	// if non empty, only listed paths will be watched within branch
	Paths []string
}

func main() {
	ymlBytes, err := ioutil.ReadFile("watch.yml")
	if err != nil {
		fail(err)
	}

	config := &WatchConfig{}

	err = yaml.Unmarshal([]byte(ymlBytes), config)
	if err != nil {
		fail(err)
	}

	fmt.Println(config)
}

func fail(e error) {
	log.Fatalf("error: %v", e)
}
