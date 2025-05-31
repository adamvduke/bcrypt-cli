// Package main implements a command-line interface for bcrypt operations.
package main

import (
	"os"
	"strconv"

	"github.com/alecthomas/kingpin/v2"

	"golang.org/x/crypto/bcrypt"

	"github.com/adamvduke/bcrypt-cli/bcryptio"
)

//go:generate go tool -modfile=tools.mod golangci-lint run --fix

const (
	version = "0.0.8"
)

const (
	appName  = "bcrypt-cli"
	appUsage = "Wraps golang.org/x/crypto/bcrypt in a cli"

	cmpCmd   = "compare"
	cmpUsage = "Compare a previously hashed password to a plain text password"

	costCmd   = "cost"
	costUsage = "Print the hashing cost used to create the given hash"

	hashCmd   = "hash"
	hashUsage = "Use bcrypt to hash a password"

	generateCmd   = "generate"
	generateUsage = "Output a random password and its bcrypt hash"

	costFlag      = costCmd
	costFlagUsage = "The hashing cost to use"

	lengthFlag      = "length"
	lengthFlagUsage = "The length of the random password to generate"

	symbolsFlag      = "symbols"
	symbolsFlagUsage = "If the generated password should contain symbols"
)

func main() {
	app := kingpin.New(appName, appUsage)
	app.Version(version).VersionFlag.Short('v')
	app.HelpFlag.Short('h')

	// Add compare command
	cmpParams := compareParams{in: os.Stdin, out: os.Stdout}
	app.Command(cmpCmd, cmpUsage).Action(cmpParams.runCompare)

	// Add cost command
	cstParams := costParams{in: os.Stdin, out: os.Stdout}
	app.Command(costCmd, costUsage).Action(cstParams.runCost)

	// Add hash command
	hParams := hashParams{in: os.Stdin, out: os.Stdout}
	hashCmd := app.Command(hashCmd, hashUsage).Action(hParams.runHash)
	hashCmd.Flag(costFlag, costFlagUsage).Short('c').Default(defaultCost()).IntVar(&hParams.cost)

	// Add generate command
	genParams := generateParams{out: os.Stdout}
	genCmd := app.Command(generateCmd, generateUsage).Action(genParams.runGenerate)
	genCmd.Flag(costFlag, costFlagUsage).Short('c').Default(defaultCost()).IntVar(&genParams.cost)
	genCmd.Flag(lengthFlag, lengthFlagUsage).Short('l').Default(defaultLength()).IntVar(&genParams.length)
	genCmd.Flag(symbolsFlag, symbolsFlagUsage).Short('s').Default("true").BoolVar(&genParams.includeSybmols)

	// Start the application
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func defaultCost() string {
	return strconv.Itoa(bcrypt.DefaultCost)
}

func defaultLength() string {
	return strconv.Itoa(bcryptio.DefaultLength)
}
