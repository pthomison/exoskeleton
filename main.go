package main

import (
	"github.com/pthomison/exoskeleton/cmd"
	utils "github.com/pthomison/golang-utils"
)

func main() {
	utils.Check(cmd.Execute())
}
