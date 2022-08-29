package asty

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func listDir(dir, suffix string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var result []string

	for _, f := range files {
		if strings.HasSuffix(f.Name(), suffix) {
			filename := filepath.Join(dir, f.Name())
			result = append(result, filename)
		}
	}

	return result, nil
}

func copyFile(src, dst string) error {
	fmt.Println("Copying", src, "to", dst)

	source, err := os.Open(src)
	if err != nil {
		return err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	err = source.Close()
	if err != nil {
		return err
	}

	err = destination.Close()
	if err != nil {
		return err
	}

	return err
}

func getTestDataRoot() string {
	dstRoot, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(dstRoot, "testdata")
}

func copyOriginalFiles() error {
	originalTestDataRoot := filepath.Join(build.Default.GOROOT, "src", "go", "printer", "testdata")
	files, err := listDir(originalTestDataRoot, ".input")
	if err != nil {
		return err
	}

	dstRoot := getTestDataRoot()

	for _, srcPath := range files {
		filename := filepath.Base(srcPath)
		dstPath := filepath.Join(dstRoot, filename)
		err = copyFile(srcPath, dstPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func compare(fileA, fileB string) error {
	match := diffmatchpatch.New()
	contentA, err := ioutil.ReadFile(fileA)
	if err != nil {
		return err
	}
	contentB, err := ioutil.ReadFile(fileB)
	if err != nil {
		return err
	}
	diffs := match.DiffMain(string(contentA), string(contentB), true)
	patches := match.PatchMake(string(contentA), diffs)
	if len(patches) > 0 {
		return fmt.Errorf("%s", match.PatchToText(patches))
	}
	return nil
}

func TestRoundTrip(t *testing.T) {
	err := copyOriginalFiles()
	if err != nil {
		t.Fatal(err)
	}

	testDataRoot := getTestDataRoot()
	files, err := listDir(testDataRoot, ".input")
	if err != nil {
		t.Fatal(err)
	}
	for _, input := range files {
		filename := filepath.Base(input)
		t.Run(filename, func(t *testing.T) {
			err := runRoundTripForFile(input)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func runRoundTripForFile(input string) error {
	stem := strings.TrimSuffix(input, ".input")

	jsonOutput := stem + ".json"
	err := SourceToJSON(input, jsonOutput, "  ", true, true, true)
	if err != nil {
		return err
	}

	output := stem + ".output"
	err = JSONToSource(jsonOutput, output, true, true, true)
	if err != nil {
		return err
	}

	golden := stem + ".golden"
	err = Loop(input, golden, true)
	if err != nil {
		return err
	}

	err = compare(output, golden)
	if err != nil {
		return err
	}
	return nil
}
