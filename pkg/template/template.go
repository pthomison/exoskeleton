package template

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/fileutils"
	"github.com/spf13/cobra"
)

const (
	CommandName = "template"
)

type Args struct {
	Input                string
	Output               string
	VariableFile         string
	CommandLineVariables []string
}

func RegisterFlags(cmd *cobra.Command, cmdArgs *Args) {
	cmd.PersistentFlags().StringVarP(&cmdArgs.Input, "input", "i", "", "file to template from")
	cmd.PersistentFlags().StringVarP(&cmdArgs.Output, "output", "o", "", "location to output the rendered template")
	cmd.PersistentFlags().StringVarP(&cmdArgs.VariableFile, "variable-file", "f", "", "file which contains yaml to inject into the template")
	cmd.PersistentFlags().StringSliceVarP(&cmdArgs.CommandLineVariables, "var", "v", []string{}, "")

	cmd.MarkPersistentFlagRequired("input")
	cmd.MarkPersistentFlagRequired("output")

}

func Run(args *Args) {
	fmt.Printf("Templating %s with %s into %s\n", args.Input, args.VariableFile, args.Output)

	var variableData fileutils.UnstructureYamlData

	if args.VariableFile != "" {
		variableData = fileutils.ReadYamlFilepath(args.VariableFile)
	} else {
		variableData = make(fileutils.UnstructureYamlData)
	}

	for _, v := range args.CommandLineVariables {
		strs := strings.SplitN(v, "=", 2)
		variableData[strs[0]] = strs[1]
	}

	inputStat, err := os.Stat(args.Input)
	errcheck.Check(err)

	if !inputStat.IsDir() {
		TemplateFile(args.Input, args.Output, variableData, inputStat.Mode())
	} else {
		fileSystem := os.DirFS(args.Input)

		fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
			errcheck.Check(err)

			info, err := d.Info()
			errcheck.Check(err)

			inputPath := fmt.Sprintf("%s/%s", args.Input, path)
			outputPath := fmt.Sprintf("%s/%s", args.Output, path)

			if !info.IsDir() {
				TemplateFile(inputPath, outputPath, variableData, info.Mode())
			}

			return nil
		})
	}
}

func TemplateFile(infile string, outfile string, varData fileutils.UnstructureYamlData, outperm os.FileMode) {
	template, err := AddTemplateFunctions(template.New(filepath.Base(infile))).ParseFiles(infile)
	errcheck.Check(err)

	outputBuffer := &bytes.Buffer{}

	err = template.Execute(outputBuffer, varData)
	errcheck.Check(err)

	err = os.MkdirAll(filepath.Dir(outfile), 0750)
	errcheck.Check(err)

	err = os.WriteFile(outfile, outputBuffer.Bytes(), outperm)
	errcheck.Check(err)
}

func AddTemplateFunctions(t *template.Template) *template.Template {

	fm := make(template.FuncMap)

	base64Decode := func(b64 string) (string, error) {
		data, err := base64.StdEncoding.DecodeString(b64)
		return string(data), err
	}

	fm["base64Decode"] = base64Decode

	return t.Funcs(fm)
}
