package main

import (
	"time"

	git "gopkg.in/src-d/go-git.v4"
)

const (
	configFile = "./watch.yml"
)

var (
	config = &WatchConfig{}
)

func main() {

	// DEV TMP
	rmDir("./repos")

	parseYML(configFile, config)

	for _, repo := range config.Repos {

		repo.openOrInitGitRepo()

		// remotes, err := repo.Remotes()
		// if err != nil {
		// 	fail(err)
		// }

		// debug("-", remotes[0].Config().URL)
	}

	saveYML(configFile, config)

	// loop forever, looking for changes
	loop()
}

func loop() {
	for {
		for name, repoInfo := range config.Repos {
			repoDir := "./repos/" + name
			repo, err := git.PlainOpen(repoDir)
			if err != nil {
				fail(err)
			}
			// fetch
			err = repo.Fetch(&git.FetchOptions{})
			if err != nil && err != git.NoErrAlreadyUpToDate {
				fail(err)
			}
			debug("fetched", repoInfo.URL)
		}

		time.Sleep(10 * time.Second)
	}
}
