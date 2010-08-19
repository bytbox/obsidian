package compile

import (
	"log"
)

func CompilePosts() {
	log.Stdout("  Compiling posts")
}

func CompileExcerpts() {
	log.Stdout("  Compiling post excerpts")
}

func CompileTags() {
	log.Stdout("  Compiling tags")
}

func CompileCategories() {
	log.Stdout("  Compiling categories")
}

func CompileIndex() {
	log.Stdout("  Compiling index page")
}

func Compile404() {
	log.Stdout("  Compiling 404 page")
}

func CompileFull() {
	log.Stdout("  Compiling full pages")
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
