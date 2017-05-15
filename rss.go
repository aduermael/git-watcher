package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gorilla/feeds"
)

func newFeedItemFromDiff(title, desc, url string, diffs []*Diff) {

	now := time.Now()

	jsonBytes, err := ioutil.ReadFile("./history.json")

	var feed *feeds.Feed
	err = json.Unmarshal(jsonBytes, &feed)
	if err != nil {
		debug("error:", err)
		return
	}

	// build feed from scratch
	if err != nil {
		feed = &feeds.Feed{
			Title:       "Git repo watcher",
			Link:        &feeds.Link{Href: ""},
			Description: "Changes from watched Git repositories.",
			Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
			Items:       make([]*feeds.Item, 0),
		}
	}

	item := &feeds.Item{
		Title:       title,
		Link:        &feeds.Link{Href: url},
		Description: "Changes:<br>",
		Author:      &feeds.Author{Name: "Git repo watcher", Email: ""},
		Created:     now,
	}

	for _, diff := range diffs {
		item.Description += string(diff.Type) + " - " + diff.File + "<br>"
	}

	feed.Items = append([]*feeds.Item{item}, feed.Items...)

	jsonBytes, err = json.Marshal(feed)
	if err != nil {
		debug("error:", err)
		return
	}

	err = ioutil.WriteFile("./history.json", jsonBytes, 0644)
	if err != nil {
		debug("error:", err)
		return
	}

	rss, _ := feed.ToRss()
	err = ioutil.WriteFile("./rss.xml", []byte(rss), 0644)
	if err != nil {
		debug("error:", err)
		return
	}

	atom, _ := feed.ToAtom()
	err = ioutil.WriteFile("./atom.xml", []byte(atom), 0644)
	if err != nil {
		debug("error:", err)
		return
	}
}
