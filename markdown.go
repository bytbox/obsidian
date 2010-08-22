package markdown

import (
	"exec"
	"os"
)

// Format takes a markdown-formatted string and returns an equivalent
// HTML-formatted string.
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
	cmd.Stdin.Close()
	_, err = cmd.Wait(0)
	return
}
