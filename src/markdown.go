package markdown

import (
	"exec"
	"io/ioutil"
	"os"
)

// Format takes a markdown-formatted string and returns an equivalent
// HTML-formatted string.
//
// TODO write built-in markdown implementation, to avoid forking for every 
// post
func Format(md string) (html string, err os.Error) {
	cmdName, err := exec.LookPath("markdown")
	if err != nil {
		return
	}
	cmd, err := exec.Run(
		cmdName,
		[]string{},
		os.Environ(),
		".",
		exec.Pipe,
		exec.Pipe,
		exec.PassThrough,
	)
	if err != nil {
		return
	}
	cmd.Stdin.WriteString(md)
	cmd.Stdin.Close()
	b, err := ioutil.ReadAll(cmd.Stdout)
	if err != nil {
		return
	}
	html = string(b)
	err = cmd.Close()
	return
}
