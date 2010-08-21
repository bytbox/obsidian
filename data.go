package data

import (
	"template"
)

type Post struct {
	Title    string
	Category string
	Tags     []string
	Content  string
	URL      string
	CompiledFull string
	CompiledExcerpt string
}

type Tag struct {
	Name  string
	Posts []*Post
}

type Category struct {
	Name  string
	Posts []*Post
}

type Page struct {
	URL     string
	Content string
	Compiled string
}

// Global data store
var (
	Posts      = map[string]*Post{}
	Tags       = map[string]*Tag{}
	Categories = map[string]*Category{}
	Pages      = map[string]*Page{}
	Templates  = make(map[string]*template.Template)
)
