package main

import (
	"asty/asty"
	"encoding/json"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func SourceToJSON(input, output string) error {
	marshaller := asty.NewMarshaller(false)
	err := marshaller.AddFile(input)
	if err != nil {
		return err
	}

	tree, err := parser.ParseFile(marshaller.FileSet(), input, nil, 0)
	if err != nil {
		return err
	}

	node := marshaller.MarshalFile(tree)

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")
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

func JSONToSource(input, output string) error {
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

	unmarshaler := asty.NewUnmarshaller()
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

func main() {
	err := SourceToJSON("test.go", "output.json")
	if err != nil {
		panic(err)
	}
	err = JSONToSource("output-processed.json", "test-processed.go")
	if err != nil {
		panic(err)
	}
}
