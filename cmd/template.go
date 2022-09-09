package cmd

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/pthomison/errcheck"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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

type UnstructureVariableData map[interface{}]interface{}

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

	variableData := ReadVariableFile(args)

	inputStat, err := os.Stat(args.Input)
	errcheck.Check(err)

	if !inputStat.IsDir() {
		TemplateFile(args.Input, args.Output, variableData, inputStat.Mode())
	} else {
		fmt.Println("dir")
	}
}

func ReadVariableFile(args *TemplateArguments) UnstructureVariableData {
	varBytes, err := os.ReadFile(args.VariableFile)
	errcheck.Check(err)

	varData := make(UnstructureVariableData)

	err = yaml.Unmarshal(varBytes, &varData)
	errcheck.Check(err)

	return varData
}

func TemplateFile(infile string, outfile string, varData UnstructureVariableData, outperm os.FileMode) {
	template, err := template.ParseFiles(infile)
	errcheck.Check(err)

	outputBuffer := &bytes.Buffer{}

	err = template.Execute(outputBuffer, varData)
	errcheck.Check(err)

	err = os.WriteFile(outfile, outputBuffer.Bytes(), outperm)
	errcheck.Check(err)
}
