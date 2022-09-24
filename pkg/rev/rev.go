package rev

import (
	"fmt"
	"io"
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
	cmd.PersistentFlags().BoolVarP(&cmdArgs.PatchRev, "patch", "z", true, "Rev the patch version")
}

func Run(args *Args, output io.Writer) {

	byteToInt := func(b []byte) int {
		str := string(b)
		i, err := strconv.Atoi(str)
		errcheck.Check(err)

		return i
	}

	matches := versionRegex.FindSubmatch([]byte(args.Input))

	if len(matches) != 5 {
		fmt.Fprintf(output, "Input does not match SemVer Input Regex: %s\n", InputRegex)
		os.Exit(1)
	}

	prefixString := string(matches[1])

	majorInt := byteToInt(matches[2])
	minorInt := byteToInt(matches[3])
	patchInt := byteToInt(matches[4])

	if args.MajorRev {
		majorInt = majorInt + 1
		minorInt = 0
		patchInt = 0
	} else if args.MinorRev {
		minorInt = minorInt + 1
		patchInt = 0
	} else if args.PatchRev {
		patchInt = patchInt + 1
	}

	fmt.Fprintf(output, "%s%d.%d.%d\n", prefixString, majorInt, minorInt, patchInt)
}
