package commands

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

// If the given io.Reader represents a terminal based stdin, readSensitive will
// read the input without displaying it on the screen. Otherwise, it reads input
// until EOF.
func readSensitive(input io.Reader) ([]byte, error) {
	var (
		data []byte
		err  error
	)
	f, ok := input.(*os.File)
	if ok && term.IsTerminal(int(f.Fd())) {
		// this should be the code path except when running tests
		data, err = term.ReadPassword(int(f.Fd()))
	} else {
		data, err = io.ReadAll(input)
	}
	trimmed := bytes.TrimSuffix(data, []byte("\n"))

	return trimmed, err
}

func readInput(input io.Reader) ([]byte, error) {
	reader := bufio.NewReader(input)
	data, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSuffix(data, "\n")
	return []byte(trimmed), nil
}
