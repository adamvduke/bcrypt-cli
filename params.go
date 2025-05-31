package main

import (
	"io"

	"github.com/alecthomas/kingpin/v2"

	"github.com/adamvduke/bcrypt-cli/bcryptio"
)

type compareParams struct {
	in  io.Reader
	out io.Writer
}

type costParams struct {
	in  io.Reader
	out io.Writer
}

type generateParams struct {
	cost           int
	length         int
	includeSybmols bool
	out            io.Writer
}

type hashParams struct {
	cost int
	in   io.Reader
	out  io.Writer
}

func (prm *compareParams) runCompare(_ *kingpin.ParseContext) error {
	return bcryptio.Compare(prm.in, prm.out)
}

func (prm *costParams) runCost(_ *kingpin.ParseContext) error {
	return bcryptio.Cost(prm.in, prm.out)
}

func (prm *generateParams) runGenerate(_ *kingpin.ParseContext) error {
	return bcryptio.Generate(prm.out, prm.includeSybmols, prm.length, prm.cost)
}

func (prm *hashParams) runHash(_ *kingpin.ParseContext) error {
	return bcryptio.Hash(prm.in, prm.out, prm.cost)
}
