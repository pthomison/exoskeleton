package ethprice

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/pthomison/errcheck"
)

func TestEthprice(t *testing.T) {

	args := &Args{
		Coin:     "ethereum",
		Currency: "usd",
	}

	buf := &bytes.Buffer{}

	Run(args, buf)

	output, err := buf.ReadString('\n')
	errcheck.CheckTest(err, t)

	output = strings.Replace(output, "\n", "", 1)

	if output == "" {
		errString := fmt.Sprintf("ethprice should not be blank")
		errcheck.CheckTest(errors.New(errString), t)
	}

}
