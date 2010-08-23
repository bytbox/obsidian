package tidy

import (
	"exec"
	"io/ioutil"
	"os"
)

// Tidy takes an HTML string and tidys it up.
//
// TODO write built-in tidy implementation, to avoid forking for every 
// page
func Tidy(str string) (html string, err os.Error) {
	cmdName, err := exec.LookPath("tidy")
	if err != nil {
		return
	}
	cmd, err := exec.Run(
		cmdName,
		[]string{},
		os.Environ(),
		".",
		exec.Pipe,
		exec.PassThrough,
		exec.PassThrough,
	)
	if err != nil {
		return
	}
	cmd.Stdin.WriteString(str)
	cmd.Stdin.Close()
	b, err := ioutil.ReadAll(cmd.Stdout)
	if err != nil {
		return
	}
	html = string(b)
	err = cmd.Close()
	return str, err
}

