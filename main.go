package main

import (
	"asty/asty"
	"flag"
	"fmt"
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

func printUsage(fs *flag.FlagSet) {
	fmt.Print(UsageString)
	fs.PrintDefaults()
}

func printError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func main() {
	fmt.Println(os.Getwd())
	args := os.Args
	var input, output string
	var indent int
	var comments, positions, references bool
	fs := flag.NewFlagSet("asty", flag.ExitOnError)
	fs.StringVar(&input, "input", "", "input file name")
	fs.StringVar(&output, "output", "", "output file name")
	fs.IntVar(&indent, "indent", 0, "indentation level")
	fs.BoolVar(&comments, "comments", false, "include comments")
	fs.BoolVar(&positions, "positions", false, "include positions")
	fs.BoolVar(&references, "references", false, "include references to reuse nodes from multiple places")

	if len(args) < 2 {
		printUsage(fs)
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
		err := asty.SourceToJSON(input, output, indentStr, comments, positions, references)
		if err != nil {
			printError(err)
		}
	case "json2go":
		err := asty.JSONToSource(input, output, comments, positions, references)
		if err != nil {
			printError(err)
		}
	case "help":
		printUsage(fs)
		return
	default:
		fmt.Printf("unknown command: %s\n", args[1])
		printUsage(fs)
		os.Exit(1)
	}
}
