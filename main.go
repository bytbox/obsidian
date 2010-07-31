package main

import (
	"fmt"
	"http"
	"opts"
	"os"
	"path"
	//"template"
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
	readData()
	
	fmt.Fprintf(os.Stderr, "Serving!\n")
	// set up the extra servers
	http.HandleFunc("/test",TestServer)
	http.HandleFunc("/",NotFoundServer)
	// start the server
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr,"%s\n",err.String())
	}
}

// The various templates.
var (
	genTemplate string
	adminTemplate string
	indexTemplate string
	postTemplate string
	tagTemplate string
	categoryTemplate string
)

func readTemplate(name string) string {
	templateDirectory := path.Join(*blogroot,"templates")
	return readFile(path.Join(templateDirectory, name))
}

func readTemplates() {
	// read the templates
	genTemplate = readTemplate("gen.html")
	adminTemplate = readTemplate("admin.html")
	indexTemplate = readTemplate("index.html")
	postTemplate = readTemplate("post.html")
	tagTemplate = readTemplate("tag.html")
	categoryTemplate = readTemplate("category.html")
}

func readData() {
	readTemplates()
}

func TestServer(c *http.Conn, req *http.Request) {
	fmt.Fprintf(c,"Hello, world!\n")
}

func NotFoundServer(c *http.Conn, req *http.Request) {
	c.WriteHeader(404)
	fmt.Fprintf(c,"404 not found\n")
}
