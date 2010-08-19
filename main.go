package main

import (
	"fmt"
	"http"
	"io/ioutil"
	"log"
	"opts"
	"os"
	"path"
	"strings"
	"template"
	"time"
	. "./util"
)

var port = opts.Single("p", "port", "the port to use", "8080")
var blogroot = opts.Single("r",
	"blogroot",
	"the root directory for blog data",
	"/usr/share/obsidian")
var showVersion = opts.Flag("", "version", "show version information")
var verbose = opts.Flag("v", "verbose", "give verbose output")

var startTime = time.Nanoseconds()

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
	makeTags()
	makeCategories()
	compileAll()
	startServer()
}

func startServer() {
	log.Stdout("Starting server")
	// set up the extra servers
	http.HandleFunc("/", NotFoundServer)
	log.Stdout("Server started in ",
		(time.Nanoseconds()-startTime)/1000,
		" microseconds")
	// start the server
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
		panic("Could not start server")
	}
}

// The various templates.
var templates = make(map[string]*template.Template)

func readTemplate(name string) *template.Template {
	log.Stdout("  Reading template ", name)
	templatePath := path.Join(templateDir, name)
	templateText := ReadFile(templatePath)
	template, err := template.Parse(templateText, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
		os.Exit(1)
	}
	return template
}

func readTemplates() {
	// read the templates
	log.Stdout("Reading templates")
	flist, err := ioutil.ReadDir(templateDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
		panic("Couldn't read template directory!")
	}
	for _, finfo := range flist {
		fname := strings.Replace(finfo.Name, ".html", "", -1)
		templates[fname] = readTemplate(fname + ".html")
	}
}

type Post struct {
	title    string
	category string
	tags     []string
	content  string
	url      string
}

type Tag struct {
	name  string
	posts []*Post
}

type Category struct {
	name  string
	posts []*Post
}

var posts = map[string]*Post{}
var tags = map[string]*Tag{}
var categories = map[string]*Category{}

type PostVisitor struct {
	root string
}

func (v PostVisitor) VisitDir(path string, f *os.FileInfo) bool {
	return true
}

func readPost(content string, path string) *Post {
	groups := strings.Split(content, "\n\n", 2)
	metalines := strings.Split(groups[0], "\n", -1)
	post := &Post{}
	post.content = groups[1]
	post.title = metalines[0]
	for _, line := range metalines[1:] {
		fmt.Printf(line)
	}
	post.url = path
	return post
}

func (v PostVisitor) VisitFile(path string, f *os.FileInfo) {
	relPath := strings.Replace(path, v.root, "", 1)
	log.Stdout("  Reading post ", relPath)
	// read in the posts
	posts[relPath] = readPost(ReadFile(path), relPath)
}

func readPosts() {
	log.Stdout("Reading posts")
	postDir := path.Join(*blogroot, "posts")
	WalkDir(postDir, PostVisitor{postDir})
}

func makeTags() {
	log.Stdout("Analyzing tags")
	for _, post := range posts {
		for _, tagname := range post.tags {
			if _, ok := tags[tagname]; !ok {
				tags[tagname] = &Tag{
					name: tagname,
					posts: make([]*Post, 0),
				}
			}
			tag := tags[tagname]
			l := len(tag.posts)
			if l+1 > cap(tag.posts) {
				newSlice := make([]*Post, (l+1)*2)
				copy(newSlice, tag.posts)
				tag.posts = newSlice
			}
			tag.posts = tag.posts[0:l+1]
			tag.posts[l] = post
		}
	}
}

func makeCategories() {
	log.Stdout("Analyzing categories")
	for _, post := range posts {
		cname := post.category
		if _, ok := categories[cname]; !ok {
			if _, ok := categories[cname]; !ok {
				categories[cname] = &Category{
					name: cname,
					posts: make([]*Post, 0),
				}
			}
			cat := categories[cname]
			l := len(cat.posts)
			if l+1 > cap(cat.posts) {
				newSlice := make([]*Post, (l+1)*2)
				copy(newSlice, cat.posts)
				cat.posts = newSlice
			}
			cat.posts = cat.posts[0:l+1]
			cat.posts[l] = post
		}
	}
}

func compilePosts() {
	log.Stdout("  Compiling posts")
}

func compileExcerpts() {
	log.Stdout("  Compiling post excerpts")
}

func compileTags() {
	log.Stdout("  Compiling tags")
}

func compileCategories() {
	log.Stdout("  Compiling categories")
}

func compileIndex() {
	log.Stdout("  Compiling index page")
}

func compile404() {
	log.Stdout("  Compiling 404 page")
}

func compileAll() {
	log.Stdout("Compiling all")
	compilePosts()
	compileExcerpts()
	compileTags()
	compileCategories()
	compileIndex()
	compile404()
}

func NotFoundServer(c *http.Conn, req *http.Request) {
	log.Stderr("404 when serving", req.URL.String())
	c.WriteHeader(404)
	fmt.Fprintf(c, "404 not found\n")
}
