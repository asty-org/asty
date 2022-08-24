package main

import (
	"asty/asty"
	"encoding/json"
	"flag"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

func SourceToJSON(input, output string, indent string, comments, positions bool) error {
	marshaller := asty.NewMarshaller(comments, positions)
	err := marshaller.AddFile(input)
	if err != nil {
		return err
	}

	tree, err := parser.ParseFile(marshaller.FileSet(), input, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	node := marshaller.MarshalFile(tree)

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", indent)
	err = encoder.Encode(node)
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func JSONToSource(input, output string, comments, positions bool) error {
	inFile, err := os.Open(input)
	if err != nil {
		return err
	}
	var decoded asty.FileNode
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(&decoded)
	if err != nil {
		return err
	}
	err = inFile.Close()
	if err != nil {
		return err
	}

	unmarshaler := asty.NewUnmarshaller(comments, positions)
	result := unmarshaler.UnmarshalFileNode(&decoded)
	fset := &token.FileSet{}

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	err = printer.Fprint(outFile, fset, result)
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}
	return nil
}

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
	args := os.Args
	var input, output string
	var indent int
	var comments, positions bool
	fs := flag.NewFlagSet("asty", flag.ExitOnError)
	fs.StringVar(&input, "input", "", "input file name")
	fs.StringVar(&output, "output", "", "output file name")
	fs.IntVar(&indent, "indent", 0, "indentation level")
	fs.BoolVar(&comments, "comments", false, "include comments")
	fs.BoolVar(&positions, "positions", false, "include positions")

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
		err := SourceToJSON(input, output, indentStr, comments, positions)
		if err != nil {
			printError(err)
		}
	case "json2go":
		err := JSONToSource(input, output, comments, positions)
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
