package data

import (
	"template"
)

// Post represents a single blog post, with certain pre-specified meta-data.
type Post struct {
	Title    string
	Category string
	Tags     []string
	Content  string
	Meta     map[string]string
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
	Title   string
	Meta    map[string]string
	Content string
	Compiled string
}

// Global data store
var (
	Posts      = map[string]*Post{}
	Tags       = map[string]*Tag{}
	Categories = map[string]*Category{}
	Pages      = map[string]*Page{}
	Templates  = map[string]*template.Template{}
	Data       = map[string]string{}
)
