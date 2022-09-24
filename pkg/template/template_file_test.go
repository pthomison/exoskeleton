package template

import (
	"errors"
	"os"
	"testing"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/fileutils"
)

const (
	singleFileFolder              = "../../resources/singleFileTemplate/"
	SingleFileTemplateFilename    = singleFileFolder + "file.template.js"
	SingleFileVariableFilename    = singleFileFolder + "vars.yaml"
	SingleFileOutputFilename      = singleFileFolder + "file.output.js"
	SingleFileOutputCheckFilename = singleFileFolder + "file.output-check.js"
)

func TestTemplateSingleFile(t *testing.T) {

	Run(&Args{
		Input:        SingleFileTemplateFilename,
		Output:       SingleFileOutputFilename,
		VariableFile: SingleFileVariableFilename,
	})

	if !fileutils.CompareFilepaths(SingleFileOutputFilename, SingleFileOutputCheckFilename) {
		errcheck.CheckTest(errors.New("template output doesn't match the valid output"), t)
	}

	err := os.Remove(SingleFileOutputFilename)
	errcheck.CheckTest(err, t)
}
