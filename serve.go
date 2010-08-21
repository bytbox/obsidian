package serve

import (
	"fmt"
	"http"
	"log"
)

// NotFoundServer displays the 404 page
var NotFoundServer = func(c *http.Conn, req *http.Request) {
	log.Stderr("404 when serving", req.URL.String())
	c.WriteHeader(404)
	fmt.Fprintf(c, "404 not found\n")
}
