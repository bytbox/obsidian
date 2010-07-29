.PHONY: all clean install

all: obsidian

obsidian: main.6
	6l -o $@ main.6

main.6: main.go
	6g -o $@ main.go

clean:
	rm obsidian *.6
