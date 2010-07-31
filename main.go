package main

import (
	"container/vector"
	"fmt"
	"http"
	"opts"
	"os"
	"path"
	"strings"
	"template"
)

var port = opts.Option("p", "port", "the port to use", "8080")
var blogroot = opts.Option("r",
	"blogroot",
	"the root directory for blog data",
	"/usr/share/obsidian")

func main() {
	// option setup
	opts.Description("lightweight http blog server")
	// parse and handle options
	opts.Parse()

	fmt.Fprintf(os.Stderr, "Reading data...\n")
	readTemplates()
	readPosts()

	fmt.Fprintf(os.Stderr, "Compiling site...\n")

	fmt.Fprintf(os.Stderr, "Serving!\n")
	// set up the extra servers
	http.HandleFunc("/test", TestServer)
	http.HandleFunc("/", NotFoundServer)
	// start the server
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
	}
}

// The various templates.
var (
	genTemplate      *template.Template
	adminTemplate    *template.Template
	indexTemplate    *template.Template
	postTemplate     *template.Template
	excerptTemplate  *template.Template
	tagTemplate      *template.Template
	categoryTemplate *template.Template
	notFoundTemplate *template.Template
)

func readTemplate(name string) *template.Template {
	templateDirectory := path.Join(*blogroot, "templates")
	templatePath := path.Join(templateDirectory, name)
	templateText := readFile(templatePath)
	template, err := template.Parse(templateText, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
		os.Exit(1)
	}
	return template
}

func readTemplates() {
	// read the templates
	genTemplate = readTemplate("gen.html")
	adminTemplate = readTemplate("admin.html")
	indexTemplate = readTemplate("index.html")
	postTemplate = readTemplate("post.html")
	tagTemplate = readTemplate("tag.html")
	categoryTemplate = readTemplate("category.html")
	excerptTemplate = readTemplate("excerpt.html")
	notFoundTemplate = readTemplate("404.html")
}

type Post struct {
	title string
	category string
	tags vector.StringVector
	content string
}

var posts = map[string] *Post{}

type PostVisitor struct {
	root string
}

func (v PostVisitor) VisitDir(path string, f *os.FileInfo) bool {
	return true
}

func readPost(content string) *Post {
	post := &Post{}
	return post
}

func (v PostVisitor) VisitFile(path string, f *os.FileInfo) {
	relPath := strings.Replace(path, v.root, "", 1)
	// read in the posts
	posts[relPath] = readPost(readFile(path))
}

func readPosts() {
	postDir := path.Join(*blogroot, "posts")
	walkDir(postDir, PostVisitor{postDir})
}

func TestServer(c *http.Conn, req *http.Request) {
	fmt.Fprintf(c, "Hello, world!\n")
}

func NotFoundServer(c *http.Conn, req *http.Request) {
	c.WriteHeader(404)
	fmt.Fprintf(c, "404 not found\n")
}
