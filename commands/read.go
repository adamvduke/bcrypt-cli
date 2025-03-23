package commands

import (
	"bufio"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

// If the given io.Reader represents a terminal based file descriptor, readSensitive
// will read the input without displaying it on the screen. Otherwise, it reads input
// until EOF.
func readSensitive(in io.Reader) ([]byte, error) {
	f, ok := in.(*os.File)
	if ok && term.IsTerminal(int(f.Fd())) {
		// this should be the code path except when running tests
		return term.ReadPassword(int(f.Fd()))
	}
	return readInput(in)
}

func readInput(in io.Reader) ([]byte, error) {
	reader := bufio.NewReader(in)
	data, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSuffix(data, "\n")
	return []byte(trimmed), nil
}
