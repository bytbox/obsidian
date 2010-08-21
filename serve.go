package serve

import (
	"fmt"
	"http"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"path"
	
	"./data"
)

// startDataServers starts a server for every entry in data.Data
func startFileServers() {
	for url, loc := range data.Data {
		http.Handle(url, FileServer{loc})
	}
}

// startPageServers starts a server for every entry in data.Pages
func startPageServers() {
	for _, page := range data.Pages {
		http.Handle(page.URL, PageServer{page})
	}
}

// startMisc starts miscellaneous servers
//
//  NotFoundServer        serves 404 error pages
func startMisc() {
	http.HandleFunc("/", NotFoundServer)
}

func StartServers() {
	log.Stdout("Starting servers")
	startFileServers()
	startPageServers()
	startMisc()
}

func Serve(port string) {
	// start the server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
		panic("Could not start server")
	}
}

// NotFoundServer displays the 404 page
var NotFoundServer = func(c *http.Conn, req *http.Request) {
	log.Stderr("404 when serving", req.URL.String())
	c.WriteHeader(404)
	fmt.Fprint(c, "404 not found\n")
}

type PageServer struct {
	Page *data.Page
}

func (p PageServer) ServeHTTP(c *http.Conn, req *http.Request) {
	fmt.Fprint(c, p.Page.Compiled)
}

type FileServer struct {
	loc string
}

func (s FileServer) ServeHTTP(c *http.Conn, req *http.Request) {
	content, err := ioutil.ReadFile(s.loc)
	if err != nil {
		log.Stderr(err)
	} else {
		// get mime type
		mimeType := mime.TypeByExtension(path.Ext(s.loc))
		c.SetHeader("Content-Type", mimeType)
		fmt.Fprintf(c, "%s", content)
	}
}
