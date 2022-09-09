package cmd

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/pthomison/errcheck"
)

const (
	templateOutput = "console.log(\"alpha + beta\");"
)

func TestTemplate(t *testing.T) {
	input := "../examples/cool-code-file.template.js"
	output := "../examples/cool-code.output"
	variableFile := "../examples/variables.yaml"

	err := os.Remove(output)

	var pathError *fs.PathError
	if (err != nil) && (!errors.As(err, &pathError)) {
		errcheck.CheckTest(err, t)
	}

	Run(&TemplateArguments{
		Input:        input,
		Output:       output,
		VariableFile: variableFile,
	})

	outputBytes, err := os.ReadFile(output)
	errcheck.CheckTest(err, t)

	if string(outputBytes) != templateOutput {
		errcheck.CheckTest(errors.New("template output doesn't match the valid output"), t)

	}

	err = os.Remove(output)
	errcheck.CheckTest(err, t)
}
