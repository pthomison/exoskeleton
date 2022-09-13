package cmd

import (
	"errors"
	"os"
	"testing"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/fileutils"
)

const (
	baseFolder            = "../resources/nestedFolderTemplate/"
	TemplateFoldername    = baseFolder + "template-folder"
	VariableFilename      = baseFolder + "vars.yaml"
	OutputFoldername      = baseFolder + "output-folder"
	OutputCheckFoldername = baseFolder + "check-folder"
)

func TestTemplateNestedFolder(t *testing.T) {
	Run(&TemplateArguments{
		Input:        TemplateFoldername,
		Output:       OutputFoldername,
		VariableFile: VariableFilename,
	})

	if fileutils.CompareFilepaths(TemplateFoldername, OutputCheckFoldername) {
		errcheck.CheckTest(errors.New("template output doesn't match the valid output"), t)
	}

	err := os.RemoveAll(OutputFoldername)
	errcheck.CheckTest(err, t)
}
