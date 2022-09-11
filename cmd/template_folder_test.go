package cmd

import (
	"errors"
	"os"
	"testing"

	"github.com/pthomison/errcheck"
)

const (
	singleFolder                      = "../resources/singleFolderTemplate/"
	SingleFolderTemplateFoldername    = singleFolder + "template-folder"
	SingleFolderVariableFilename      = singleFolder + "vars.yaml"
	SingleFolderOutputFoldername      = singleFolder + "output-folder"
	SingleFolderOutputCheckFoldername = singleFolder + "check-folder"
)

func TestTemplateSingleFolder(t *testing.T) {

	Run(&TemplateArguments{
		Input:        SingleFolderTemplateFoldername,
		Output:       SingleFolderOutputFoldername,
		VariableFile: SingleFolderVariableFilename,
	})

	// outputBytes, err := os.ReadFile(singleFileTemplateOutputName)
	// errcheck.CheckTest(err, t)

	if CompareFolders(SingleFolderOutputFoldername, SingleFolderOutputCheckFoldername) {
		errcheck.CheckTest(errors.New("template output doesn't match the valid output"), t)
	}

	err := os.RemoveAll(SingleFolderOutputFoldername)
	errcheck.CheckTest(err, t)
}