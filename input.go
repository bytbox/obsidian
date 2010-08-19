package input

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"template"
	
	. "./data"
)

func readDir(dirname string) (ret []*os.FileInfo) {
	ret, err := ioutil.ReadDir(dirname)
	if err != nil {
		// the directory doesn't exist - create it
		err = os.MkdirAll(dirname, 0755)
	}
	return
}

func walkDir(dirname string, v path.Visitor) {
	// first make sure the directory exists
	readDir(dirname)
	// now walk it
	path.Walk(dirname, v, nil)
}

func readFile(filename string) (contents string) {
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


func ReadTemplate(templateDir, name string) *template.Template {
	log.Stdout("  Reading template ", name)
	templatePath := path.Join(templateDir, name)
	templateText := readFile(templatePath)
	template, err := template.Parse(templateText, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
		os.Exit(1)
	}
	return template
}

func ReadTemplates(templateDir string) (templates map[string]*template.Template) {
	templates = make(map[string]*template.Template)
	// read the templates
	log.Stdout("Reading templates")
	flist, err := ioutil.ReadDir(templateDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
		panic("Couldn't read template directory!")
	}
	for _, finfo := range flist {
		fname := strings.Replace(finfo.Name, ".html", "", -1)
		templates[fname] = ReadTemplate(templateDir, fname + ".html")
	}
	return
}

func ReadPost(content string, path string) *Post {
	groups := strings.Split(content, "\n\n", 2)
	metalines := strings.Split(groups[0], "\n", -1)
	post := &Post{}
	post.Content = groups[1]
	post.Title = metalines[0]
	for _, line := range metalines[1:] {
		fmt.Printf(line)
	}
	post.URL = path
	return post
}

type postVisitor struct {
	root string
	posts map[string]*Post
}

func (v postVisitor) VisitDir(path string, f *os.FileInfo) bool {
	return true
}

func (v postVisitor) VisitFile(path string, f *os.FileInfo) {
	relPath := strings.Replace(path, v.root, "", 1)
	log.Stdout("  Reading post ", relPath)
	// read in the posts
	v.posts[relPath] = ReadPost(readFile(path), relPath)
}

func ReadPosts(postDir string) map[string]*Post {
	log.Stdout("Reading posts")
	v := postVisitor{root: postDir, posts: make(map[string]*Post)}
	walkDir(postDir, v)
	return v.posts
}

