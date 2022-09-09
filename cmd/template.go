package cmd

import (
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
		Run:   run,
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

func run(cmd *cobra.Command, args []string) {
	fmt.Printf("Templating %s with %s into %s\n", templateArgs.Input, templateArgs.VariableFile, templateArgs.Output)

	fmt.Println("4")

	// templateInfo, err := os.Stat(templateArgs.Input)
	// errcheck.Check(err)

	fmt.Println("3")

	template, err := template.ParseGlob(templateArgs.Input)
	errcheck.Check(err)

	fmt.Println("1")

	varBytes, err := os.ReadFile(templateArgs.VariableFile)
	errcheck.Check(err)

	fmt.Println("2")

	varData := make(map[interface{}]interface{})

	err = yaml.Unmarshal(varBytes, &varData)
	errcheck.Check(err)

	// var templateBytes bytes.Buffer

	err = template.Execute(os.Stdout, varData)
	errcheck.Check(err)

	// err = os.WriteFile(templateArgs.Output, templateBytes.Bytes(), templateInfo.Mode().Perm())
	// errcheck.Check(err)

}
