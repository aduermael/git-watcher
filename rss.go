package main

import (
	"time"

	"github.com/gorilla/feeds"
)

func diffToRSSFeed(title, desc string, diffs []*Diff) {

	now := time.Now()

	feed := &feeds.Feed{
		Title:       "Git repo watcher",
		Description: "Changes from watched Git repositories.",
		Created:     now,
		Items:       make([]*feeds.Item, 0),
	}

	item := &feeds.Item{
		Title:       "changes",
		Link:        &feeds.Link{Href: "https://github.com"},
		Description: "Changes:\n",
		Created:     now,
	}

	for _, diff := range diffs {
		item.Description += string(diff.Type) + " - " + diff.File + "\n"
	}

	feed.Items = append(feed.Items, item)

	rss, _ := feed.ToRss()

	debug(rss)
}
