package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	message string
	name    string

	templateCmd = &cobra.Command{
		Use:   "template",
		Short: "Template Utility",
		Long:  `template dir + yaml vars = magic`,
		Run:   run,
	}
)

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.PersistentFlags().StringVarP(&message, "message", "m", "hello world", "message the program will output")
	templateCmd.PersistentFlags().StringVarP(&name, "name", "n", "patrick", "name the program will output to")

}

func run(cmd *cobra.Command, args []string) {
	fmt.Println(message + ", " + name)
}
