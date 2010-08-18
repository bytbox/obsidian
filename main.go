package main

import (
	"container/vector"
	"fmt"
	"http"
	"io/ioutil"
	"opts"
	"os"
	"path"
	"strings"
	"template"
)

var port = opts.Single("p", "port", "the port to use", "8080")
var blogroot = opts.Single("r",
	"blogroot",
	"the root directory for blog data",
	"/usr/share/obsidian")
var showVersion = opts.Flag("", "version", "show version information")
var verbose = opts.Flag("v", "verbose", "give verbose output")

var (
	templateDir string
)

func main() {
	// option setup
	opts.Description = "lightweight http blog server"
	// parse and handle options
	opts.Parse()
	
	templateDir = path.Join(*blogroot, "templates")
	
	readTemplates()
	readPosts()
	
	// set up the extra servers
	http.HandleFunc("/", NotFoundServer)
	// start the server
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
	}
}

// The various templates.
var templates = make(map[string]*template.Template)

func readTemplate(name string) *template.Template {
	templatePath := path.Join(templateDir, name)
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
	flist, err := ioutil.ReadDir(templateDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
		panic("Couldn't read template directory!")
	}
	for _, finfo := range flist {
		fname := strings.Replace(finfo.Name, ".html", "", -1)
		templates[fname] = readTemplate(fname+".html")
	}
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

func NotFoundServer(c *http.Conn, req *http.Request) {
	c.WriteHeader(404)
	fmt.Fprintf(c, "404 not found\n")
}
