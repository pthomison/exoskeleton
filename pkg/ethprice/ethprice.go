package ethprice

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pthomison/errcheck"
	"github.com/spf13/cobra"
	coingecko "github.com/superoo7/go-gecko/v3"
)

const (
	CommandName = "ethprice"
)

type Args struct {
	Coin     string
	Currency string
}

func RegisterFlags(cmd *cobra.Command, cmdArgs *Args) {
	cmd.PersistentFlags().StringVarP(&cmdArgs.Coin, "coin", "", "ethereum", "coin to price")
	cmd.PersistentFlags().StringVarP(&cmdArgs.Currency, "currency", "", "usd", "currency to price in")
}

func Run(args *Args, output io.Writer) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	ids := []string{args.Coin}
	vc := []string{args.Currency}
	sp, err := CG.SimplePrice(ids, vc)
	errcheck.Check(err)

	price := (*sp)[args.Coin][args.Currency]

	fmt.Fprintf(output, "%f\n", price)
}
