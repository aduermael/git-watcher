package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func parseYML(ymlPath string, conf *WatchConfig) {
	ymlBytes, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		fail(err)
	}

	err = yaml.Unmarshal(ymlBytes, config)
	if err != nil {
		fail(err)
	}

	for repoKey, repo := range config.Repos {
		repo.Name = repoKey
		for branchKey, branch := range repo.Branches {
			if branch == nil {
				repo.Branches[branchKey] = &Branch{Name: branchKey}
			} else {
				branch.Name = branchKey
			}
		}
	}
}

func saveYML(ymlPath string, conf *WatchConfig) {
	ymlBytes, err := yaml.Marshal(conf)
	if err != nil {
		fail(err)
	}
	ioutil.WriteFile(ymlPath, ymlBytes, 0644)
}
