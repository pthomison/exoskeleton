# exoskeleton

## golang powered CLI tool for personal utilities

```
$ exoskeleton --help

Usage:
  exoskeleton [flags]
  exoskeleton [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  template    Template Utility

Flags:
  -h, --help   help for exoskeleton

Use "exoskeleton [command] --help" for more information about a command.
```

```
$ exoskeleton template --help

template dir + yaml vars = magic

Usage:
  exoskeleton template [flags]

Flags:
  -h, --help                   help for template
  -i, --input string           file to template from
  -o, --output string          location to output the rendered template
  -f, --variable-file string   file which contains yaml to inject into the template
```