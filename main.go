package main

import (
	"github.com/pthomison/errcheck"
	"github.com/pthomison/exoskeleton/cmd"
)

func main() {
	errcheck.Check(cmd.Execute())
}
