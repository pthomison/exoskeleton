package cmd

import (
	"github.com/pthomison/ssm-k8s-injector/pkg/injector"
	ssmk8sinjector "github.com/pthomison/ssm-k8s-injector/pkg/injector"
	"github.com/spf13/cobra"
)

var (
	ssmK8sInjectorArguments = &injector.Args{}

	ssmK8sInjectorCmd = &cobra.Command{
		Use:   ssmk8sinjector.CommandName,
		Short: ssmk8sinjector.CommandName,
		Run:   ssmK8sInjectorRun,
	}
)

func init() {
	rootCmd.AddCommand(ssmK8sInjectorCmd)
	injector.RegisterFlags(ssmK8sInjectorCmd, ssmK8sInjectorArguments)
}

func ssmK8sInjectorRun(cmd *cobra.Command, args []string) {
	injector.Run(ssmK8sInjectorArguments)
}
