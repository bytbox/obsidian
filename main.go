package main

import (
	"log"
	"opts"
	"path"
	"time"

	compile "./compile"
	.       "./data"
	input   "./input"
	serve   "./serve"
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
	postDir     string
	pageDir     string
	dataDir     string
)

func main() {
	// option setup
	opts.Description = "lightweight http blog server"
	// parse and handle options
	opts.Parse()

	templateDir = path.Join(*blogroot, "templates")
	postDir = path.Join(*blogroot, "posts")
	pageDir = path.Join(*blogroot, "pages")
	dataDir = path.Join(*blogroot, "data")

	input.ReadTemplates(templateDir)
	input.ReadPosts(postDir)
	input.ReadPages(pageDir)
	input.ReadData(dataDir)
	makeTags()
	makeCategories()
	compile.CompileAll()
	serve.StartServers()
	log.Stdout("Server started in ",
		(time.Nanoseconds()-startTime)/1000,
		" microseconds")
	serve.Serve(*port)
}

func makeTags() {
	log.Stdout("Analyzing tags")
	for _, post := range Posts {
		for _, tagname := range post.Tags {
			if _, ok := Tags[tagname]; !ok {
				Tags[tagname] = &Tag{
					Name:  tagname,
					Posts: make([]*Post, 0),
				}
			}
			tag := Tags[tagname]
			l := len(tag.Posts)
			if l+1 > cap(tag.Posts) {
				newSlice := make([]*Post, (l+1)*2)
				copy(newSlice, tag.Posts)
				tag.Posts = newSlice
			}
			tag.Posts = tag.Posts[0 : l+1]
			tag.Posts[l] = post
		}
	}
}

func makeCategories() {
	log.Stdout("Analyzing categories")
	for _, post := range Posts {
		cname := post.Category
		if _, ok := Categories[cname]; !ok {
			if _, ok := Categories[cname]; !ok {
				Categories[cname] = &Category{
					Name:  cname,
					Posts: make([]*Post, 0),
				}
			}
			cat := Categories[cname]
			l := len(cat.Posts)
			if l+1 > cap(cat.Posts) {
				newSlice := make([]*Post, (l+1)*2)
				copy(newSlice, cat.Posts)
				cat.Posts = newSlice
			}
			cat.Posts = cat.Posts[0 : l+1]
			cat.Posts[l] = post
		}
	}
}
