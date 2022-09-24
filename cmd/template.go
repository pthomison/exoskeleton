package cmd

import (
	"github.com/pthomison/exoskeleton/pkg/template"
	"github.com/spf13/cobra"
)

var (
	templateArguments = &template.Args{}

	templateCmd = &cobra.Command{
		Use:   template.CommandName,
		Short: template.CommandName,
		Run:   templateRun,
	}
)

func init() {
	rootCmd.AddCommand(templateCmd)
	template.RegisterFlags(templateCmd, templateArguments)
}

func templateRun(cmd *cobra.Command, args []string) {
	template.Run(templateArguments)
}
