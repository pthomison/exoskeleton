package cmd

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pthomison/errcheck"
	"github.com/pthomison/fileutils"
	"github.com/spf13/cobra"
)

var (
	templateArgs = &TemplateArguments{}

	templateCmd = &cobra.Command{
		Use:   "template",
		Short: "Template Utility",
		Long:  `template dir + yaml vars = magic`,
		Run:   cobraRun,
	}
)

type TemplateArguments struct {
	Input        string
	Output       string
	VariableFile string
}

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.PersistentFlags().StringVarP(&templateArgs.Input, "input", "i", "", "file to template from")
	templateCmd.PersistentFlags().StringVarP(&templateArgs.Output, "output", "o", "", "location to output the rendered template")
	templateCmd.PersistentFlags().StringVarP(&templateArgs.VariableFile, "variable-file", "f", "", "file which contains yaml to inject into the template")

}

func cobraRun(cmd *cobra.Command, args []string) {
	Run(templateArgs)
}

func Run(args *TemplateArguments) {
	fmt.Printf("Templating %s with %s into %s\n", args.Input, args.VariableFile, args.Output)

	variableData := fileutils.ReadYamlFilepath(args.VariableFile)

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
	template, err := template.ParseFiles(infile)
	errcheck.Check(err)

	outputBuffer := &bytes.Buffer{}

	err = template.Execute(outputBuffer, varData)
	errcheck.Check(err)

	err = os.MkdirAll(filepath.Dir(outfile), 0750)
	errcheck.Check(err)

	err = os.WriteFile(outfile, outputBuffer.Bytes(), outperm)
	errcheck.Check(err)
}
