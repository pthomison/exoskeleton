package getparam

import (
	"fmt"
	"io"

	"github.com/pthomison/awsutils"
	"github.com/pthomison/errcheck"
	"github.com/spf13/cobra"
)

const (
	CommandName = "getparam"
)

type Args struct {
	SSMparameter string
	AWSRegion    string
}

func RegisterFlags(cmd *cobra.Command, cmdArgs *Args) {
	// cmd.PersistentFlags().StringVarP(&cmdArgs.SSMparameter, "ssm-parameter", "p", "", "Parameter to inject")
	cmd.PersistentFlags().StringVarP(&cmdArgs.AWSRegion, "aws-region", "r", "us-east-2", "")

	// cmd.MarkPersistentFlagRequired("ssm-parameter")
}

func Run(parameterPath string, args *Args, output io.Writer) {
	value, err := awsutils.AWSGetParameter(parameterPath, args.AWSRegion)
	errcheck.Check(err)

	fmt.Fprintf(output, "%v\n", value)

	// cs, err := k8sutils.GetClientSet()
	// errcheck.Check(err)

	// trueSecretKey := ""

	// if args.SecretKey == "" {
	// 	trueSecretKey = filepath.Base(args.SSMparameter)
	// } else {
	// 	trueSecretKey = args.SecretKey
	// }

	// sec := make(k8sutils.Secret)
	// sec[trueSecretKey] = []byte(value)

	// _, err = k8sutils.GetSecret(cs, args.SecretName, args.K8Snamespace)

	// if err != nil {
	// 	_, err = k8sutils.ApplySecret(cs, args.SecretName, args.K8Snamespace, sec)
	// 	errcheck.Check(err)
	// } else {
	// 	_, err = k8sutils.UpdateSecret(cs, args.SecretName, args.K8Snamespace, sec)
	// 	errcheck.Check(err)
	// }
}
