package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	git "gopkg.in/src-d/go-git.v4"
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

// Branch contains what needs to be watched in the git branch
// If no specific filters are set, then any change will be reported
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

	for name, repoInfo := range config.Repos {

		repoDir := "./repos/" + name

		repo, err := git.PlainOpen(repoDir)

		// repo found and opened, but is it the one we're looking for?
		// we should check the URL and delete directory if it doesn't match
		for err == nil {
			var remotes []*git.Remote
			remotes, err = repo.Remotes()

			if err != nil {
				rmDir(repoDir)
				break
			}

			// TODO: allow several remotes
			if len(remotes) > 1 {
				err = errors.New("only one remote expected")
				rmDir(repoDir)
				break
			}

			if repoInfo.URL != remotes[0].Config().URL {
				err = errors.New("remote URL is different from the one in the config")
				debug(err)
				rmDir(repoDir)
				break
			}

			break
		}

		if err != nil {
			// if the repo does not exist already, clone it
			if err == git.ErrRepositoryNotExists {
				repo, err = git.PlainClone("./repos/"+name, true, &git.CloneOptions{URL: repoInfo.URL /*, Depth: 1*/})
			}
			if err != nil {
				fail(err)
			}
		}

		remotes, err := repo.Remotes()
		if err != nil {
			fail(err)
		}

		debug("-", remotes[0].Config().URL)
	}
}

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
