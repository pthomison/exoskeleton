package rev

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/pthomison/errcheck"
	"github.com/spf13/cobra"
)

const (
	CommandName = "rev"

	InputRegex = `^([vV]?)(\d+)\.(\d+)\.(\d+)$`
)

var (
	versionRegex = regexp.MustCompile(InputRegex)
)

type Args struct {
	Input    string
	MajorRev bool
	MinorRev bool
	PatchRev bool
}

func RegisterFlags(cmd *cobra.Command, cmdArgs *Args) {
	cmd.PersistentFlags().StringVarP(&cmdArgs.Input, "input", "i", "", "file to template from")
	cmd.PersistentFlags().BoolVarP(&cmdArgs.MajorRev, "major", "x", false, "Rev the major version")
	cmd.PersistentFlags().BoolVarP(&cmdArgs.MinorRev, "minor", "y", false, "Rev the minor version")
	cmd.PersistentFlags().BoolVarP(&cmdArgs.PatchRev, "patch", "z", false, "Rev the patch version")
}

func Run(args *Args) {
	matches := versionRegex.FindSubmatch([]byte(args.Input))

	if len(matches) != 5 {
		fmt.Println("Input does not match SemVer Input Regex")
		os.Exit(1)
	}

	prefixString := string(matches[1])

	majorString := string(matches[2])
	majorInt, err := strconv.Atoi(majorString)
	errcheck.Check(err)

	minorString := string(matches[3])
	minorInt, err := strconv.Atoi(minorString)
	errcheck.Check(err)

	patchString := string(matches[4])
	patchInt, err := strconv.Atoi(patchString)
	errcheck.Check(err)

	if args.MajorRev {
		majorInt = majorInt + 1
	}

	if args.MinorRev {
		minorInt = minorInt + 1
	}

	if args.PatchRev {
		patchInt = patchInt + 1
	}

	fmt.Printf("%s%d.%d.%d\n", prefixString, majorInt, minorInt, patchInt)
}
