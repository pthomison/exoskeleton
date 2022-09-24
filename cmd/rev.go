package cmd

import (
	"os"

	"github.com/pthomison/exoskeleton/pkg/rev"
	"github.com/spf13/cobra"
)

var (
	revArguments = &rev.Args{}

	revCmd = &cobra.Command{
		Use:   rev.CommandName,
		Short: rev.CommandName,
		Run:   revRun,
	}
)

func init() {
	rootCmd.AddCommand(revCmd)
	rev.RegisterFlags(revCmd, revArguments)
}

func revRun(cmd *cobra.Command, args []string) {
	rev.Run(revArguments, os.Stdout)
}
