.PHONY: all clean install

include ${GOROOT}/src/Make.${GOARCH}

all: obsidian

obsidian: main.${O}
	${LD} -o $@ main.${O}

main.${O}: main.go util.go
	${GC} -o $@ main.go util.go

clean:
	rm obsidian *.${O}
