.PHONY: all clean install

include ${GOROOT}/src/Make.${GOARCH}

all: obsidian

obsidian: main.${O}
	${LD} -o $@ main.${O}

main.${O}: main.go
	${GC} -o $@ main.go

clean:
	rm obsidian *.${O}
