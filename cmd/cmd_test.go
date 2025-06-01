package cmd_test

import (
	"context"
	"errors"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"golang.org/x/crypto/bcrypt"

	clireader "github.com/adamvduke/cli-reader"

	"github.com/adamvduke/bcrypt-cli/cmd"
)

const testVersion = "0.0.999"

func TestCompare(t *testing.T) {
	cases := []struct {
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
	for _, tsc := range cases {
		t.Run(tsc.name, func(t *testing.T) {
			args := []string{"compare"}
			stdin := clireader.New(tsc.hashed, tsc.plain)
			stdout := &strings.Builder{}
			root := stubbedRoot(stdin, stdout, args)
			ctx := context.Background()

			if err := root.ExecuteContext(ctx); err != nil {
				if !errors.Is(err, tsc.err) {
					t.Fatal(err)
				}
			}
			got := stdout.String()
			if !strings.Contains(got, tsc.want) {
				t.Errorf("Expected output to contain %q, got %q", tsc.want, got)
			}
		})
	}
}

func TestCost(t *testing.T) {
	cases := []struct {
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
	for _, tsc := range cases {
		t.Run(tsc.name, func(t *testing.T) {
			args := []string{"cost"}
			stdin := clireader.New(tsc.hashed)
			stdout := &strings.Builder{}
			root := stubbedRoot(stdin, stdout, args)
			ctx := context.Background()

			if err := root.ExecuteContext(ctx); err != nil {
				if !errors.Is(err, tsc.err) {
					t.Fatal(err)
				}
			}
			got := stdout.String()
			if got != tsc.want {
				t.Errorf("\ngot: \n%s\nwant: \n%s", got, tsc.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	cases := []struct {
		name   string
		cost   string
		length string
		outLen int
		err    error
	}{
		{
			name:   "cost 12 length 25",
			cost:   "12",
			length: "25",
			outLen: 25,
		},
		{
			name:   "below min length",
			cost:   "10",
			length: "1",
			outLen: 20,
		},
		{
			name:   "above max length",
			cost:   "10",
			length: "73",
			outLen: 72,
		},
	}

	for _, tsc := range cases {
		t.Run(tsc.name, func(t *testing.T) {
			args := []string{"generate", "--length", tsc.length, "--symbols", "--cost", tsc.cost}
			stdin := clireader.New("") // No input needed for generate command
			stdout := &strings.Builder{}
			root := stubbedRoot(stdin, stdout, args)
			ctx := context.Background()

			if err := root.ExecuteContext(ctx); err != nil {
				if !errors.Is(err, tsc.err) {
					t.Fatal(err)
				}
			}

			// Plaintext: USEUZDLHRQX7KMRJTDLXU6BRZ7LN64ABWIFLSYNGAM2P5T2SD7
			// Bcrypt: $2a$12$3idUylIztAwQkUvcw2AQeOsL2HzfrMvy3IRumEKA74HnJJ75fedHa
			raw := strings.TrimSpace(stdout.String())
			lines := strings.Split(raw, "\n")
			if len(lines) != 2 {
				t.Fatal("incorrect number of output lines", len(lines))
			}
			pw := strings.TrimSpace(strings.Split(lines[0], " ")[1])
			if len(pw) != tsc.outLen {
				t.Errorf("incorrect length got: %d, want: %d", len(pw), tsc.outLen)
			}

			bc := strings.TrimSpace(strings.Split(lines[1], " ")[1])
			cost, err := bcrypt.Cost([]byte(bc))
			if err != nil {
				t.Fatal("unexpected error", err)
			}
			if strconv.Itoa(cost) != tsc.cost {
				t.Errorf("incorrect cost got: %d, want: %s", cost, tsc.cost)
			}
		})
	}
}

func TestHash(t *testing.T) {
	cases := []struct {
		name  string
		plain string
		cost  string
		want  string
		err   error
	}{
		{
			name:  "cost 12",
			plain: "ZoO19Tx}p;eVM)fv.!2=",
			cost:  "12",
		},
	}
	for _, tsc := range cases {
		t.Run(tsc.name, func(t *testing.T) {
			args := []string{"hash", "--cost", tsc.cost}
			stdin := clireader.New(tsc.plain, tsc.plain)
			stdout := &strings.Builder{}
			root := stubbedRoot(stdin, stdout, args)
			ctx := context.Background()

			if err := root.ExecuteContext(ctx); err != nil {
				if !errors.Is(err, tsc.err) {
					t.Fatal(err)
				}
			}
			raw := strings.TrimSpace(stdout.String())
			lines := strings.Split(raw, "\n")
			if len(lines) != 3 {
				t.Fatal("incorrect number of output lines", len(lines))
			}

			bc := strings.TrimSpace(strings.Split(lines[2], " ")[1])
			cost, err := bcrypt.Cost([]byte(bc))
			if err != nil {
				t.Fatal("unexpected error", err)
			}
			if strconv.Itoa(cost) != tsc.cost {
				t.Errorf("incorrect cost got: %d, want: %s", cost, tsc.cost)
			}
		})
	}
}

func TestVersionCommand(t *testing.T) {
	cases := []string{"version", "--version"}
	for _, tsc := range cases {
		t.Run(tsc, func(t *testing.T) {
			args := []string{tsc}
			stdin := clireader.New()
			stdout := &strings.Builder{}
			root := stubbedRoot(stdin, stdout, args)
			ctx := context.Background()

			if err := root.ExecuteContext(ctx); err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			got := stdout.String()
			want := "bcrypt-cli version 0.0.999\n"
			if got != want {
				t.Errorf("Expected output to contain %q, got %q", want, got)
			}
		})
	}
}

func stubbedRoot(in io.Reader, out io.Writer, args []string) *cobra.Command {
	root := cmd.NewRootCmd(testVersion)
	root.SetArgs(args)
	root.SetIn(in)
	root.SetOut(out)
	return root
}
