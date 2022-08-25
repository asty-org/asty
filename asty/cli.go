package asty

import (
	"encoding/json"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func SourceToJSON(input, output string, indent string, comments, positions, references bool) error {
	marshaller := NewMarshaller(comments, positions, references)

	mode := parser.SkipObjectResolution
	if comments {
		mode |= parser.ParseComments
	}
	tree, err := parser.ParseFile(marshaller.FileSet(), input, nil, mode)
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

func JSONToSource(input, output string, comments, positions, references bool) error {
	inFile, err := os.Open(input)
	if err != nil {
		return err
	}
	var node FileNode
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(&node)
	if err != nil {
		return err
	}
	err = inFile.Close()
	if err != nil {
		return err
	}

	unmarshaler := NewUnmarshaller(comments, positions, references)
	tree := unmarshaler.UnmarshalFileNode(&node)

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	err = printer.Fprint(outFile, unmarshaler.FileSet(), tree)
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func Loop(input, output string, comments bool) error {
	mode := parser.SkipObjectResolution
	if comments {
		mode |= parser.ParseComments
	}
	fs := token.NewFileSet()
	tree, err := parser.ParseFile(fs, input, nil, mode)
	if err != nil {
		return err
	}

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	err = printer.Fprint(outFile, fs, tree)
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}
	return nil
}
