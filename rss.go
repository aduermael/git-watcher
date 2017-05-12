package main

import (
	"time"

	"github.com/gorilla/feeds"
)

func diffToRSSFeed(title, desc string, diffs []*Diff) {

	now := time.Now()

	feed := &feeds.Feed{
		Title:       "Git repo watcher",
		Link:        &feeds.Link{Href: ""},
		Description: "Changes from watched Git repositories.",
		Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
		Created:     now,
		Items:       make([]*feeds.Item, 0),
	}

	item := &feeds.Item{
		Title:       "changes",
		Link:        &feeds.Link{Href: "https://github.com"},
		Description: "Changes:<br>",
		Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
		Created:     now,
	}

	for _, diff := range diffs {
		item.Description += string(diff.Type) + " - " + diff.File + "<br>"
	}

	feed.Items = append(feed.Items, item)

	rss, _ := feed.ToRss()

	debug(rss)
}
