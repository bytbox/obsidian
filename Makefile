.PHONY: all clean install

all: obsidian

include ${GOROOT}/src/Make.${GOARCH}

.SUFFIXES: .go .${O}

obsidian: main.${O}
	${LD} -o $@ main.${O}

.go.${O}:
	${GC} -o $@ $*.go

format:
	gofmt -w ${GOFILES}

clean:
	rm -f obsidian *.${O}

include Makefile.deps
