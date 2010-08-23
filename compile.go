package compile

import (
	"log"
	"os"
	"strings"

	config "./config"
	. "./data"
	tidy "./tidy"
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
			Title:   post.Title,
			Content: w.buff,
		}
	}
}

func CompileExcerpts() {
	log.Stdout("  Compiling post excerpts")
	// compile against the "excerpt" template
	tmpl := Templates["excerpt"]
	for _, post := range Posts {
		// extract excerpt
		post.Excerpt = strings.Split(post.Content, "<!--more-->", 2)[0]
		w := &stringWriter{}
		tmpl.Execute(post, w)
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
	}
}

func CompileIndex() {
	log.Stdout("  Compiling index page")
	// compile list of all pages against the "index" template
	tmpl := Templates["index"]
	w := &stringWriter{}
	tmpl.Execute(map[string]interface{}{
		"Posts": Posts,
	}, w)
	Pages["/"] = &Page {
		URL:     "/",
		Content: w.buff,
	}
}

func Compile404() {
	log.Stdout("  Compiling 404 page")
	tmpl := Templates["404"]
	w := &stringWriter{}
	tmpl.Execute(map[string]interface{}{
		"Config": config.Configuration,
	}, w)
	Pages["/404"] = &Page {
		URL:     "/-", // prevent 404 page from ever being served as a valid page
		Content: w.buff,
	}
}

func CompileFull() {
	log.Stdout("  Compiling full pages")
	// Compile all pages against the "gen" template
	tmpl := Templates["gen"]
	for _, page := range Pages {
		w := &stringWriter{}
		tmpl.Execute(map[string]interface{}{
			"Page": page,
			"Config": config.Configuration,
			"Pages": Pages,
			"Tags": Tags,
			"Categories": Categories,
		}, w)
		page.Compiled, _ = tidy.Tidy(w.buff)
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
