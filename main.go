package main

import "time"

const (
	configFile  = "/data/watch.yml"
	historyFile = "/data/history.json"
	rssFile     = "/data/rss.xml"
	atomFile    = "/data/atom.xml"
	reposDir    = "/data/repos"
)

var (
	config = &WatchConfig{}
)

func main() {
	parseYML(configFile, config)

	for _, repo := range config.Repos {
		repo.openOrInitGitRepo()
	}

	// if new repositories have been initiated, it's a good time to save
	// commits in the configuration file.
	saveYML(configFile, config)

	// http server to expose index, rss and atom
	go serveHTTP()

	// loop forever, looking for changes
	loop()
}

func loop() {
	for {
		for _, repo := range config.Repos {
			// compare commits, looking for changes
			// TODO: error handling
			err := repo.fetchAndLookForChanges()
			if err != nil {
				debug("error:", err.Error())
			}
		}

		saveYML(configFile, config)

		time.Sleep(1 * time.Hour)
	}
}
