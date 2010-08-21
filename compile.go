package compile

import (
	"log"
	"os"

	. "./data"
)

// string writer is an io.Writer implementation that just writes to a stirng
type stringWriter struct {
	buff string
}

func (w *stringWriter) Write(data []uint8) (n int, err os.Error) {
	str := string(data)
	n = len(str)
	w.buff += str
	return
}

func CompilePosts() {
	log.Stdout("  Compiling posts")
	// compile against the "post" template
	tmpl := Templates["post"]
	for _, post := range Posts {
		w := &stringWriter{}
		tmpl.Execute(post, w)
		post.CompiledFull = w.buff
		Pages[post.URL] = &Page{
			URL:     post.URL,
			Content: w.buff,
		}
	}
}

func CompileExcerpts() {
	log.Stdout("  Compiling post excerpts")
	// compile against the "excerpt" template
	tmpl := Templates["excerpt"]
	for _, post := range Posts {
		w := &stringWriter{}
		tmpl.Execute(post, w) // TODO get excerpt
		post.CompiledExcerpt = w.buff
	}
}

func CompileTags() {
	log.Stdout("  Compiling tags")
	// compile against the "tag" template
	tmpl := Templates["tab"]
	for _, tag := range Tags {
		w := &stringWriter{}
		tmpl.Execute(tag, w)
		Pages["/tag/"+tag.Name] = &Page {
			URL:     "/tag/"+tag.Name,
			Content: w.buff,
		}
	}
}

func CompileCategories() {
	log.Stdout("  Compiling categories")
	// compile against the "tag" template
	tmpl := Templates["category"]
	for _, cat := range Categories {
		w := &stringWriter{}
		tmpl.Execute(cat, w)
		Pages["/category/"+cat.Name] = &Page {
			URL:     "/category/"+cat.Name,
			Content: w.buff,
		}
	}}

func CompileIndex() {
	log.Stdout("  Compiling index page")
	// TODO
}

func Compile404() {
	log.Stdout("  Compiling 404 page")
	// TODO
}

func CompileFull() {
	log.Stdout("  Compiling full pages")
	tmpl := Templates["gen"]
	for _, page := range Pages {
		w := &stringWriter{}
		tmpl.Execute(page, w)
		page.Compiled = w.buff
	}
}

func CompileAll() {
	log.Stdout("Compiling all")
	CompilePosts()
	CompileExcerpts()
	CompileTags()
	CompileCategories()
	CompileIndex()
	Compile404()
	CompileFull()
}
