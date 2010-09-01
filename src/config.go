package config

import (
	"log"
	"io/ioutil"
	"strings"
)

type Config map[string]interface{}

var Configuration Config

func asList(val string) (list []string) {
	list = strings.Split(val, ",", -1)
	for i, str := range list {
		list[i] = strings.TrimSpace(str)
	}
	return
}

func ReadConfig(confFile string) {
	log.Stdout("Reading configuration")
	content, _ := ioutil.ReadFile(confFile)
	lines := strings.Split(string(content), "\n", -1)
	Configuration = make(map[string]interface{})
	for _, line := range lines {
		ind := strings.Index(line, "=")
		if ind != -1 {
			key, value := strings.TrimSpace(line[0:ind]),
				strings.TrimSpace(line[ind+1:])
			Configuration[strings.Title(key)] = value
			
			// and as a list
			listkey := strings.Title(key)+"List"
			Configuration[listkey] = asList(value)
		}
	}
}
