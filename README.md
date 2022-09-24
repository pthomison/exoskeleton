# exoskeleton

## golang powered CLI tool for personal utilities

### Install Instructions
```
brew tap pthomison/homebrew-tools
brew install pthomison/tools/exoskeleton
```

### Top Level

```
exoskeleton

Usage:
  exoskeleton [flags]
  exoskeleton [command]

Available Commands:
  completion       Generate the autocompletion script for the specified shell
  ethprice         ethprice
  help             Help about any command
  rev              rev
  ssm-k8s-injector ssm-k8s-injector
  template         template

Flags:
  -h, --help   help for exoskeleton

Use "exoskeleton [command] --help" for more information about a command.

```


### Templating Tool

```
template

Usage:
  exoskeleton template [flags]

Flags:
  -h, --help                   help for template
  -i, --input string           file to template from
  -o, --output string          location to output the rendered template
  -v, --var strings            
  -f, --variable-file string   file which contains yaml to inject into the template

```

### AWS SSM -> Kubernetes Secret Injection Tool

```
ssm-k8s-injector

Usage:
  exoskeleton ssm-k8s-injector [flags]

Flags:
  -r, --aws-region string       (default "us-east-2")
  -h, --help                   help for ssm-k8s-injector
  -n, --k8s-namespace string   Namespace of the secret (default "default")
  -k, --secret-key string      Secret to inject into
  -s, --secret-name string     Secret to inject into
  -p, --ssm-parameter string   Parameter to inject

```

### Semantic Version Versioning Tool

```
rev

Usage:
  exoskeleton rev [flags]

Flags:
  -h, --help           help for rev
  -i, --input string   file to template from
  -x, --major          Rev the major version
  -y, --minor          Rev the minor version
  -z, --patch          Rev the patch version (default true)

```

### CryptoCurrency Pricing Tool (via CoinGecko APIs)

```
ethprice

Usage:
  exoskeleton ethprice [flags]

Flags:
      --coin string       coin to price (default "ethereum")
      --currency string   currency to price in (default "usd")
  -h, --help              help for ethprice

```