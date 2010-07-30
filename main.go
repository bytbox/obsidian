package main

import (
	"opts"
)

var port = opts.Option("p", "port", "the port to use", "8080")
var blogroot = opts.Option("r", 
	"blogroot", 
	"the root directory for blog data",
	"/var/obsidian")

func main() {
	opts.Description("lightweight http blog server")
	opts.Parse()
}
