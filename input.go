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
	markdown "./markdown"
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

func ReadTemplates(templateDir string) {
	Templates = make(map[string]*template.Template)
	// read the templates
	log.Stdout("Reading templates")
	flist, err := ioutil.ReadDir(templateDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
		panic("Couldn't read template directory!")
	}
	for _, finfo := range flist {
		fname := strings.Replace(finfo.Name, ".html", "", -1)
		Templates[fname] = ReadTemplate(templateDir, fname+".html")
	}
}

func ReadPost(content string, path string) *Post {
	groups := strings.Split(content, "\n\n", 2)
	metalines := strings.Split(groups[0], "\n", -1)
	post := &Post{}
	post.Content, _ = markdown.Format(groups[1])
	post.Title = metalines[0]
	post.Meta = make(map[string]string)
	for _, line := range metalines[1:] {
		ind := strings.Index(line, ":")
		if ind != -1 {
			key, value := line[0:ind], strings.TrimSpace(line[ind+1:])
			post.Meta[strings.Title(key)] = value
		}
	}
	post.URL = path
	
	// clean the post
	if len(post.Meta["category"]) == 0 {
		post.Category = "General" // TODO make this configurable
	} else {
		post.Category = post.Meta["category"]
	}
	return post
}

// postVisitor is used to havigate the directory of posts and create posts
type postVisitor struct {
	root  string
	posts map[string]*Post
}

func (v postVisitor) VisitDir(path string, f *os.FileInfo) bool {
	return true
}

func (v postVisitor) VisitFile(path string, f *os.FileInfo) {
	// get a clean path
	relPath := strings.Replace(path, v.root, "", 1)
	log.Stdout("  Reading post ", relPath)
	// read in the posts
	v.posts[relPath] = ReadPost(readFile(path), relPath)
}

// ReadPosts reads all the posts from the given directory
func ReadPosts(postDir string) {
	log.Stdout("Reading posts")
	v := postVisitor{root: postDir, posts: make(map[string]*Post)}
	walkDir(postDir, v)
	Posts = v.posts
}

// pageVisitor is used to havigate the directory of posts and create posts
type pageVisitor struct {
	root  string
}

func (v pageVisitor) VisitDir(path string, f *os.FileInfo) bool { return true }

func (v pageVisitor) VisitFile(path string, f *os.FileInfo) {
	// get a clean path
	relPath := strings.Replace(path, v.root, "", 1)
	log.Stdout("  Reading page ", relPath)
	// read in the posts
	Pages[relPath] = ReadPage(readFile(path), relPath)
}

func ReadPage(content string, path string) *Page {
	groups := strings.Split(content, "\n\n", 2)
	metalines := strings.Split(groups[0], "\n", -1)
	page := &Page{}
	page.Content, _ = markdown.Format(groups[1])
	page.Title = metalines[0]
	page.Meta = make(map[string]string)
	for _, line := range metalines[1:] {
		ind := strings.Index(line, ":")
		if ind != -1 {
			key, value := line[0:ind], strings.TrimSpace(line[ind+1:])
			page.Meta[strings.Title(key)] = value
		}
	}
	page.URL = path
	return page
}

// ReadPages reads all raw pages from the given directory
func ReadPages(pageDir string) {
	log.Stdout("Reading pages")
	v := pageVisitor{pageDir}
	walkDir(pageDir, v)
}


// dataVisitor is used to havigate the directory of data
type dataVisitor struct {
	root  string
}

func (v dataVisitor) VisitDir(path string, f *os.FileInfo) bool {
	return true
}

func (v dataVisitor) VisitFile(path string, f *os.FileInfo) {
	// get a clean path
	relPath := strings.Replace(path, v.root, "", 1)
	// note the location of this datum
	Data[relPath] = path
}

// ReadData reads all raw data from the given directory
func ReadData(dataDir string) {
	log.Stdout("Reading data")
	v := dataVisitor{root: dataDir}
	walkDir(dataDir, v)
}
