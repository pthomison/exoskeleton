package cmd

// import (
// 	"errors"
// 	"os"
// 	"testing"

// 	"github.com/pthomison/errcheck"
// )

// const (
// 	singleFolderFolder             = "tmp"
// 	singleFolderName               = "/example"
// 	singleFolderTemplateAlphaName  = singleFolderFolder + singleFolderName + "/fileAlpha.tpl"
// 	singleFolderTemplateBetaName   = singleFolderFolder + singleFolderName + "/fileBeta.tpl"
// 	singleFolderVariableFileName   = singleFolderFolder + singleFolderName + "/vars.yaml"
// 	singleFolderTemplateOutputName = singleFolderFolder + singleFolderName + "/file.output"

// 	singleFileTemplateAlpha = `console.log("{{.variableA}} + {{.variableB}}");`
// 	singleFileTemplateBeta  = `console.log("{{.variableC}} + {{.variableD}}");`

// 	singleFileVariableFile = `---
// variableA: "alpha"
// variableB: "beta"
// variableC: "charlie"
// variableD: "beta"
// `

// 	singleFileTemplateOutput = `console.log(\"alpha + beta\");`

// 	singleFilePerms = 0750
// )

// func TestTemplateSingleFile(t *testing.T) {

// 	err := os.Mkdir(singleFileFolder, singleFilePerms)
// 	errcheck.CheckTest(err, t)

// 	err = os.WriteFile(singleFileTemplateName, []byte(singleFileTemplate), singleFilePerms)
// 	errcheck.CheckTest(err, t)

// 	err = os.WriteFile(singleFileVariableFileName, []byte(singleFileVariableFile), singleFilePerms)
// 	errcheck.CheckTest(err, t)

// 	Run(&TemplateArguments{
// 		Input:        singleFileTemplateName,
// 		Output:       singleFileTemplateOutputName,
// 		VariableFile: singleFileVariableFileName,
// 	})

// 	outputBytes, err := os.ReadFile(singleFileTemplateOutputName)
// 	errcheck.CheckTest(err, t)

// 	if string(outputBytes) != templateOutput {
// 		errcheck.CheckTest(errors.New("template output doesn't match the valid output"), t)

// 	}

// 	err = os.RemoveAll(singleFileFolder)
// 	errcheck.CheckTest(err, t)
// }
