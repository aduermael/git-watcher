package main

import "time"

const (
	configFile = "./watch.yml"
)

var (
	config = &WatchConfig{}
)

func main() {

	parseYML(configFile, config)

	for _, repo := range config.Repos {
		repo.openOrInitGitRepo()
	}

	saveYML(configFile, config)

	// loop forever, looking for changes
	loop()
}

func loop() {
	for {
		for _, repo := range config.Repos {
			// compare commits, looking for changes
			_ = repo.fetchAndLookForChanges()
		}

		time.Sleep(10 * time.Second)
	}
}
