package asty

import (
	"encoding/json"
	"go/parser"
	"go/printer"
	"go/token"
)

type Options struct {
	WithPositions  bool
	WithComments   bool
	WithReferences bool
	WithImports    bool
}

func SourceToJSON(input, output string, indent string, options Options) error {
	marshaller := NewMarshaller(options)

	mode := parser.SkipObjectResolution
	if options.WithComments {
		mode |= parser.ParseComments
	}

	inFile, closeIn, err := OpenRead(input)
	if err != nil {
		return err
	}
	defer closeIn()

	tree, err := parser.ParseFile(marshaller.FileSet(), input, inFile, mode)
	if err != nil {
		return err
	}

	node := marshaller.MarshalFile(tree)

	outFile, closeOut, err := OpenOrCreateWrite(output)
	if err != nil {
		return err
	}
	defer closeOut()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", indent)
	err = encoder.Encode(node)
	if err != nil {
		return err
	}
	return nil
}

func JSONToSource(input, output string, options Options) error {
	inFile, closeIn, err := OpenRead(input)
	if err != nil {
		return err
	}
	defer closeIn()

	var node FileNode
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(&node)
	if err != nil {
		return err
	}

	unmarshaler := NewUnmarshaller(options)
	tree := unmarshaler.UnmarshalFileNode(&node)

	outFile, closeOut, err := OpenOrCreateWrite(output)
	if err != nil {
		return err
	}
	defer closeOut()
	err = printer.Fprint(outFile, unmarshaler.FileSet(), tree)
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

	outFile, closeOut, err := OpenOrCreateWrite(output)
	if err != nil {
		return err
	}
	defer closeOut()

	err = printer.Fprint(outFile, fs, tree)
	if err != nil {
		return err
	}
	return nil
}
