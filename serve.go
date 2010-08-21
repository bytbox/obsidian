package serve

import (
	"fmt"
	"http"
	"log"
	"os"
	
	"./data"
)

// startDataServers starts a server for every entry in data.Data
func startDataServers() {
	
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
	startDataServers()
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
