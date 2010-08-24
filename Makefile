.PHONY: all clean install format

all: obsidian

include ${GOROOT}/src/Make.${GOARCH}
include Makefile.info

.SUFFIXES: .go .${O} .a

obsidian: src/main.${O}
	${LD} -o $@ src/main.${O}

.go.${O}:
	${GC} -o $*.${O} $*.go

.go.a:
	${GC} -o $*.${O} $*.go && gopack grc $*.a $*.${O}

format:
	gofmt -w ${GOFILES}

cleanpackages:
	rm -f ${GOPACKAGES}

cleanarchives:
	rm -f ${GOARCHIVES}

clean: cleanpackages cleanarchives
	rm -f obsidian

include Makefile.deps
