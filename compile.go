package compile

import (
	"log"
	"os"

	. "./data"
)

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
	// TODO
}

func CompileCategories() {
	log.Stdout("  Compiling categories")
	// TODO
}

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
		tmpl.Execute(page.Content, w)
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
