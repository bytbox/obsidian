package data

type Post struct {
	Title    string
	Category string
	Tags     []string
	Content  string
	URL      string
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
	URL string
	Content string
}
