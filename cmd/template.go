package cmd

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/pthomison/errcheck"
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

	variableData := readYamlFile(args.VariableFile)

	inputStat, err := os.Stat(args.Input)
	errcheck.Check(err)

	if !inputStat.IsDir() {
		TemplateFile(args.Input, args.Output, variableData, inputStat.Mode())
	} else {
		fmt.Println("dir")

		dir, err := os.Open(args.Input)
		errcheck.Check(err)

		recurseDir(dir, func(leafNode *os.File, outputLocation string, variableData UnstructureYamlData) {
			return
		})

	}
}

func TemplateFile(infile string, outfile string, varData UnstructureYamlData, outperm os.FileMode) {
	template, err := template.ParseFiles(infile)
	errcheck.Check(err)

	outputBuffer := &bytes.Buffer{}

	err = template.Execute(outputBuffer, varData)
	errcheck.Check(err)

	err = os.WriteFile(outfile, outputBuffer.Bytes(), outperm)
	errcheck.Check(err)
}
