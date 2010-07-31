package main

import (
	"fmt"
	"http"
	"io/ioutil"
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

func ReadDir(dirname string) (ret []*os.FileInfo) {
	ret, err := ioutil.ReadDir(dirname)
	if err != nil {
		// the directory doesn't exist - create it
		err = os.MkdirAll(dirname, 0755)
	}
	return
}

func ReadFile(filename string) (contents string) {
	contentarry, err := ioutil.ReadFile(filename)
	if err != nil {
		// the file doesn't exist - create the directory and the file
		dirname, _ := path.Split(filename)
		err = os.MkdirAll(dirname, 0755)
		ioutil.WriteFile(filename, []byte{}, 0644)
	}
	contents = string(contentarry)
	return
}

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
	return ReadFile(path.Join(templateDirectory, name))
}

func readData() {
	// read the templates
	genTemplate = readTemplate("gen.html")
	adminTemplate = readTemplate("admin.html")
	indexTemplate = readTemplate("index.html")
	postTemplate = readTemplate("post.html")
	tagTemplate = readTemplate("tag.html")
	categoryTemplate = readTemplate("category.html")
}

func TestServer(c *http.Conn, req *http.Request) {
	fmt.Fprintf(c,"Hello, world!\n")
}

func NotFoundServer(c *http.Conn, req *http.Request) {
	c.WriteHeader(404)
	fmt.Fprintf(c,"404 not found\n")
}
