package template

import (
	"errors"
	"os"
	"testing"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/fileutils"
)

const (
	singleFolder                      = "./resources/singleFolderTemplate/"
	SingleFolderTemplateFoldername    = singleFolder + "template-folder"
	SingleFolderVariableFilename      = singleFolder + "vars.yaml"
	SingleFolderOutputFoldername      = singleFolder + "output-folder"
	SingleFolderOutputCheckFoldername = singleFolder + "check-folder"
)

func TestTemplateSingleFolder(t *testing.T) {
	Run(&Args{
		Input:        SingleFolderTemplateFoldername,
		Output:       SingleFolderOutputFoldername,
		VariableFile: SingleFolderVariableFilename,
	})

	if !fileutils.CompareFilepaths(SingleFolderOutputFoldername, SingleFolderOutputCheckFoldername) {
		errcheck.CheckTest(errors.New("template output doesn't match the valid output"), t)
	}

	err := os.RemoveAll(SingleFolderOutputFoldername)
	errcheck.CheckTest(err, t)
}
