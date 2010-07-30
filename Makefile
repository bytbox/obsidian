.PHONY: all clean install

include ${GOROOT}/src/Make.${GOARCH}

all: obsidian

obsidian: main.${O}
	${O}l -o $@ main.${O}

main.${O}: main.go
	${O}g -o $@ main.go

clean:
	rm obsidian *.${O}
