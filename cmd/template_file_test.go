package cmd

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/pthomison/errcheck"
)

const (
	singleFileFolder              = "../resources/singleFileTemplate/"
	SingleFileTemplateFilename    = singleFileFolder + "file.template.js"
	SingleFileVariableFilename    = singleFileFolder + "vars.yaml"
	SingleFileOutputFilename      = singleFileFolder + "file.output.js"
	SingleFileOutputCheckFilename = singleFileFolder + "file.output-check.js"
)

func TestTemplateSingleFile(t *testing.T) {

	Run(&TemplateArguments{
		Input:        SingleFileTemplateFilename,
		Output:       SingleFileOutputFilename,
		VariableFile: SingleFileVariableFilename,
	})

	outputBytes, err := os.ReadFile(SingleFileOutputFilename)
	errcheck.CheckTest(err, t)

	outputCheckBytes, err := os.ReadFile(SingleFileOutputCheckFilename)
	errcheck.CheckTest(err, t)

	if bytes.Compare(outputBytes, outputCheckBytes) != 0 {
		errcheck.CheckTest(errors.New("template output doesn't match the valid output"), t)
	}

	err = os.Remove(SingleFileOutputFilename)
	errcheck.CheckTest(err, t)
}
