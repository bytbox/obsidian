.PHONY: all clean install

all: obsidian

include ${GOROOT}/src/Make.${GOARCH}
include Makefile.info

.SUFFIXES: .go .${O}

obsidian: main.${O}
	${LD} -o $@ main.${O}

.go.${O}:
	${GC} $*.go

.go.a:
	${GC} -o $*.${O} $*.go
	gopack grc $*.a $*.${O}
	rm $*.${O}

format:
	gofmt -w ${GOFILES}

clean:
	rm -f obsidian *.a *.${O}

include Makefile.deps
