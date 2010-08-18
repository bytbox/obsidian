.PHONY: all clean install

include ${GOROOT}/src/Make.${GOARCH}

FILES = main.go util.go

all: obsidian

obsidian: main.${O}
	${LD} -o $@ main.${O}

main.${O}: ${FILES}
	${GC} -o $@ ${FILES}

format:
	gofmt -w ${FILES}

clean:
	rm obsidian *.${O}
