package main

import (
	"flag"
	"fmt"
	"github.com/asty-org/asty/asty"
	"os"
	"strings"
)

const UsageString = `Usage: asty <command> [flags]
commands:
  go2json - convert go source to json
  json2go - convert json to go source
  help    - print this message
flags:
`

func printError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func main() {
	args := os.Args
	var input, output string
	var indent int
	var comments, positions, references, imports bool
	fs := flag.NewFlagSet("asty", flag.ExitOnError)
	fs.StringVar(&input, "input", "", "input file name (default: stdin)")
	fs.StringVar(&output, "output", "", "output file name (default: stdout)")
	fs.IntVar(&indent, "indent", 0, "indentation level (default: 0)")
	fs.BoolVar(&comments, "comments", false, "include comments (default: false)")
	fs.BoolVar(&positions, "positions", false, "include positions (default: false)")
	fs.BoolVar(&references, "references", false,
		"include references to reuse nodes from multiple places (default: false)")
	fs.BoolVar(&imports, "imports", false,
		"include imports list into output (default: false)")

	fs.Usage = func() {
		fmt.Fprint(fs.Output(), UsageString)
		fs.PrintDefaults()
	}

	if len(args) < 2 {
		fs.Usage()
		return
	}

	err := fs.Parse(args[2:])
	if err != nil {
		printError(err)
	}

	if input == "" {
		input = os.Stdin.Name()
	}

	if output == "" {
		output = os.Stdout.Name()
	}

	switch args[1] {
	case "go2json":
		indentStr := strings.Repeat(" ", indent)
		options := asty.Options{
			WithImports:    imports,
			WithComments:   comments,
			WithPositions:  positions,
			WithReferences: references,
		}
		err := asty.SourceToJSON(input, output, indentStr, options)
		if err != nil {
			printError(err)
		}
	case "json2go":
		err := asty.JSONToSource(input, output, comments, positions, references)
		if err != nil {
			printError(err)
		}
	case "help":
		fs.Usage()
		return
	default:
		fmt.Printf("unknown command: %s\n", args[1])
		fs.Usage()
		os.Exit(1)
	}
}
