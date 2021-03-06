package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gorilla/feeds"
)

func serveHTTP() {
	http.HandleFunc("/", index) // `/` matches all paths not matched
	http.HandleFunc("/rss", rss)
	http.HandleFunc("/feed", rss)
	http.HandleFunc("/atom", atom)
	fail(http.ListenAndServe(":80", nil))
}

var (
	htmlTemplate *template.Template
)

func index(w http.ResponseWriter, r *http.Request) {
	var err error

	if htmlTemplate == nil {
		htmlTemplate, err = template.ParseFiles("./index.tmpl")
		if err != nil {
			http.Error(w, "can't parse template", 500)
			return
		}
	}

	var jsonBytes []byte
	jsonBytes, err = ioutil.ReadFile(historyFile)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var feed *feeds.Feed
	err = json.Unmarshal(jsonBytes, &feed)
	if err != nil {
		debug("error:", err)
		return
	}

	htmlTemplate.Execute(w, feed)
}

func rss(w http.ResponseWriter, r *http.Request) {
	rssBytes, err := ioutil.ReadFile(rssFile)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(rssBytes)
}

func atom(w http.ResponseWriter, r *http.Request) {
	rssBytes, err := ioutil.ReadFile(atomFile)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(rssBytes)
}
