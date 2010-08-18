.PHONY: all clean install

include ${GOROOT}/src/Make.${GOARCH}

.SUFFIXES: .go .${O}

all: obsidian

include Makefile.deps

obsidian: main.${O}
	${LD} -o $@ main.${O}

.go.${O}:
	${GC} -o $@ $*.go

format:
	gofmt -w *.go

clean:
	rm -f obsidian *.${O}
