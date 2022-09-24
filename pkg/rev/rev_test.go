package rev

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/pthomison/errcheck"
)

func RevTest(t *testing.T, args *Args, validOutput string) {
	buf := &bytes.Buffer{}

	Run(args, buf)

	output, err := buf.ReadString('\n')
	errcheck.CheckTest(err, t)

	output = strings.Replace(output, "\n", "", 1)

	if output != validOutput {
		errString := fmt.Sprintf("rev output doesn't match validation answers (%s vs %s)", validOutput, output)
		errcheck.CheckTest(errors.New(errString), t)
	}
}

func TestMajorRev(t *testing.T) {

	args := &Args{
		Input:    "v0.1.1",
		MajorRev: true,
		MinorRev: false,
		PatchRev: false,
	}

	validOutput := "v1.0.0"

	RevTest(t, args, validOutput)
}

func TestMinorRev(t *testing.T) {

	args := &Args{
		Input:    "v0.0.1",
		MajorRev: false,
		MinorRev: true,
		PatchRev: false,
	}

	validOutput := "v0.1.0"

	RevTest(t, args, validOutput)
}

func TestPatchRev(t *testing.T) {

	args := &Args{
		Input:    "v0.0.1",
		MajorRev: false,
		MinorRev: false,
		PatchRev: true,
	}

	validOutput := "v0.0.2"

	RevTest(t, args, validOutput)
}
