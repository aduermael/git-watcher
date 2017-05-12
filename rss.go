package main

type Channel struct {
	Title       string
	Link        string
	Description string
	Items       []*Item
}

type Item struct {
	Title       string
	Link        string
	Description string
}
