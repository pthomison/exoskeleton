package cmd

import (
	"os"

	"github.com/pthomison/exoskeleton/pkg/ethprice"
	"github.com/spf13/cobra"
)

var (
	ethpriceArguments = &ethprice.Args{}

	ethpriceCmd = &cobra.Command{
		Use:   ethprice.CommandName,
		Short: ethprice.CommandName,
		Run:   ethpriceRun,
	}
)

func init() {
	rootCmd.AddCommand(ethpriceCmd)
	ethprice.RegisterFlags(ethpriceCmd, ethpriceArguments)
}

func ethpriceRun(cmd *cobra.Command, args []string) {
	ethprice.Run(ethpriceArguments, os.Stdout)
}
