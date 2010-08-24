package config

import (
	"log"
	"io/ioutil"
	"strings"
)

type Config map[string]string

var Configuration Config

func ReadConfig(confFile string) {
	log.Stdout("Reading configuration")
	content, _ := ioutil.ReadFile(confFile)
	lines := strings.Split(string(content), "\n", -1)
	Configuration = make(map[string]string)
	for _, line := range lines {
		ind := strings.Index(line, "=")
		if ind != -1 {
			key, value := strings.TrimSpace(line[0:ind]), 
				strings.TrimSpace(line[ind+1:])
			Configuration[strings.Title(key)] = value
		}
	}
}
