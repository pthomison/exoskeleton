package cmd

import (
	"os"

	"github.com/pthomison/exoskeleton/pkg/getparam"
	"github.com/spf13/cobra"
)

var (
	getparamArguments = &getparam.Args{}

	getparamCmd = &cobra.Command{
		Use:   "getparam [parameter path]",
		Short: getparam.CommandName,
		Run:   getparamRun,
		Args:  cobra.ExactArgs(1),
	}
)

func init() {
	rootCmd.AddCommand(getparamCmd)
	getparam.RegisterFlags(getparamCmd, getparamArguments)
}

func getparamRun(cmd *cobra.Command, args []string) {
	getparam.Run(args[0], getparamArguments, os.Stdout)
}
