package commands_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/adamvduke/bcrypt-cli/commands"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/crypto/bcrypt"

	clireader "github.com/adamvduke/cli-reader"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		name   string
		plain  string
		hashed string
		want   string
		err    error
	}{
		{
			name:   "valid match",
			plain:  "ZoO19Tx}p;eVM)fv.!2=",
			hashed: "$2a$10$iouZNNkxdwPTZvo0GiOxY.2G6JTrk4VGmQH1lIe3kDQLA/E.RCTQ6",
			want:   "Enter previously hashed password:\nEnter password:\nPassword is correct\n",
			err:    nil,
		},
		{
			name:   "short hash",
			plain:  "ZoO19Tx}p;eVM)fv.!2=",
			hashed: "$2a$10$iouZNN",
			want:   "Enter previously hashed password:\nEnter password:\n",
			err:    bcrypt.ErrHashTooShort,
		},
		{
			name:   "empty plain text",
			plain:  "",
			hashed: "$2a$10$iouZNNkxdwPTZvo0GiOxY.2G6JTrk4VGmQH1lIe3kDQLA/E.RCTQ6",
			want:   "Enter previously hashed password:\nEnter password:\n",
			err:    bcrypt.ErrMismatchedHashAndPassword,
		},
		{
			name:   "incorrect hash",
			plain:  "ZoO19Tx}p;eVM)fv.!2=",
			hashed: "$2a$10$iouZNNkxdwPTZvo0GiOxY.NOPENOPENOPENOPENOPENOPENOPENOPE",
			want:   "Enter previously hashed password:\nEnter password:\n",
			err:    bcrypt.ErrMismatchedHashAndPassword,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := clireader.New(test.hashed, test.plain)
			writer := &strings.Builder{}
			c := &commands.CompareCommand{In: reader, Out: writer}
			if err := c.Run(nil); err != nil {
				if !errors.Is(err, test.err) {
					t.Fatal(err)
				}
			}
			got := writer.String()
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("got: \n%s, want: \n%s, \ndiff:\n %s", got, test.want, diff)
			}
		})
	}
}

func TestCost(t *testing.T) {
	tests := []struct {
		name   string
		hashed string
		want   string
		err    error
	}{
		{
			name:   "10",
			hashed: "$2a$10$iouZNNkxdwPTZvo0GiOxY.2G6JTrk4VGmQH1lIe3kDQLA/E.RCTQ6",
			want:   "Enter previously hashed password:\nCost: 10\n",
		},
		{
			name:   "15",
			hashed: "$2a$15$pUUuP0mNP.8X729GPd04x.25sGWFN3ZajSIZc.01MXHwkSY/HxUTi",
			want:   "Enter previously hashed password:\nCost: 15\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := clireader.New(test.hashed)
			writer := &strings.Builder{}
			c := &commands.CostCommand{In: reader, Out: writer}
			if err := c.Run(nil); err != nil {
				if !errors.Is(err, test.err) {
					t.Fatal(err)
				}
			}
			got := writer.String()
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("got: \n%s, want: \n%s, \ndiff:\n %s", got, test.want, diff)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		name            string
		cost            int
		requestedLength int
		generatedLength int
		err             error
	}{
		{
			name:            "cost 12 length 25",
			cost:            12,
			requestedLength: 25,
			generatedLength: 25,
		},
		{
			name:            "below min length",
			cost:            10,
			requestedLength: 1,
			generatedLength: 20,
		},
		{
			name:            "above max length",
			cost:            10,
			requestedLength: 73,
			generatedLength: 72,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := &strings.Builder{}
			c := &commands.GenerateCommand{Out: writer, Length: test.requestedLength, Cost: test.cost}
			if err := c.Run(nil); err != nil {
				if !errors.Is(err, test.err) {
					t.Fatal(err)
				}
			}
			raw := writer.String()
			lines := strings.Split(raw, "\n")
			if len(lines) != 3 {
				t.Fatal("incorrect number of output lines", len(lines))
			}
			pw := strings.TrimSpace(strings.Split(lines[0], " ")[1])
			if len(pw) != test.generatedLength {
				t.Errorf("incorrect length got: %d, want: %d", len(pw), test.generatedLength)
			}

			bc := strings.TrimSpace(strings.Split(lines[1], " ")[1])
			cost, err := bcrypt.Cost([]byte(bc))
			if err != nil {
				t.Fatal("unexpected error", err)
			}
			if cost != test.cost {
				t.Errorf("incorrect cost got: %d, want: %d", cost, test.cost)
			}
		})
	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		name  string
		plain string
		cost  int
		want  string
		err   error
	}{
		{
			name:  "cost 12",
			plain: "ZoO19Tx}p;eVM)fv.!2=",
			cost:  12,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := clireader.New(test.plain, test.plain)
			writer := &strings.Builder{}
			c := &commands.HashCommand{In: reader, Out: writer, Cost: test.cost}
			if err := c.Run(nil); err != nil {
				if !errors.Is(err, test.err) {
					t.Fatal(err)
				}
			}
			raw := writer.String()
			lines := strings.Split(raw, "\n")
			if len(lines) != 4 {
				t.Fatal("incorrect number of output lines", len(lines))
			}

			bc := strings.TrimSpace(strings.Split(lines[2], " ")[1])
			cost, err := bcrypt.Cost([]byte(bc))
			if err != nil {
				t.Fatal("unexpected error", err)
			}
			if cost != test.cost {
				t.Errorf("incorrect cost got: %d, want: %d", cost, test.cost)
			}
		})
	}
}
