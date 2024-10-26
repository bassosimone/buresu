// SPDX-License-Identifier: GPL-3.0-or-later

// Package txtartesting provides utilities for testing code
// where test cases are stored in .txtar files.
//
// The expected format of a .txtar file is as follows:
//
// 1. Each test case is stored in a separate file.
// 2. The file name is the test case name.
// 3. The file content must contain three sections:
//
//	-- input --
//	-- output --
//	-- error --
//
// The input section contains the input to the code under test.
//
// The output section contains the expected output.
//
// The error section contains the expected error message.
//
// The output and error sections are mutually exclusive but at
// least one of them must be present.
//
// No other files are allowed in the .txtar archive.
//
// The output section should contain either the expected output
// from executing the code under test or a serialized JSON.
//
// You load a test case by calling [LoadTestCases] with the path to
// the directory containing the .txtar files (typically, this is the
// `testdata` directory). The function returns a [TestCase] slice,
// each representing a single test case.
//
// When the expected output is text, you can compare it with the
// actual output by calling the [*TestCase.CompareTextOutput] method on
// the test case. When the expected output is JSON, you can compare it
// with the actual output by calling [*TestCase.CompareJSONOutput].
//
// When you expect an error, you can compare it with the actual
// error by calling the [*TestCase.CompareError] method.
//
// We strip leading and trailing white spaces from the input, output,
// and error sections. In particular, for output sections, we are
// careful to strip trailing newlines, removing only the empty lines.
package txtartesting

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/txtar"
)

// TestCase represents a single test case extracted from a .txtar file.
type TestCase struct {
	// Name is the file name.
	Name string

	// Input is the input to parse.
	Input string

	// Output is the expected output.
	Output string

	// Error is the expected error message.
	Error string
}

// LoadSingleTestCase loads a single `.txtar` test case from the given file path.
func LoadSingleTestCase(filePath string) (*TestCase, error) {
	archiveData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read txtar file: %v", err)
	}

	archive := txtar.Parse(archiveData)
	testCase := &TestCase{Name: filepath.Base(filePath)}

	for _, file := range archive.Files {
		switch file.Name {
		case "input":
			testCase.Input = strings.TrimSpace(string(file.Data))
		case "output":
			testCase.Output = trimEmptyLines(string(file.Data))
		case "error":
			testCase.Error = strings.TrimSpace(string(file.Data))
		default:
			return nil, fmt.Errorf("unexpected file %s in test case %s", file.Name, testCase.Name)
		}
	}

	if testCase.Output == "" && testCase.Error == "" {
		return nil, fmt.Errorf("test case %s must have either output or error", testCase.Name)
	}

	return testCase, nil
}

// LoadTestCases loads `.txtar` test cases from the given directory.
func LoadTestCases(testdataDir string) ([]*TestCase, error) {
	files, err := os.ReadDir(testdataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read testdata directory: %v", err)
	}

	var testCases []*TestCase
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".txtar" {
			continue
		}

		tc, err := LoadSingleTestCase(filepath.Join(testdataDir, file.Name()))
		if err != nil {
			return nil, err
		}
		testCases = append(testCases, tc)
	}

	return testCases, nil
}

// CompareError compares the obtained error with the expected error.
func (tc *TestCase) CompareError(err error) error {
	switch {
	case tc.Error != "" && err == nil:
		return fmt.Errorf("expected error, got none in test case %s", tc.Name)

	case tc.Error != "" && err != nil:
		if diff := cmp.Diff(tc.Error, err.Error()); diff != "" {
			return fmt.Errorf("error mismatch in test case %s (-expected +got):\n%s", tc.Name, diff)
		}
		return nil

	case tc.Error == "" && err != nil:
		return fmt.Errorf("unexpected error in test case %s: %v", tc.Name, err)

	default:
		return nil
	}
}

// CompareTextOutput compares the obtained output with the expected output.
func (tc *TestCase) CompareTextOutput(output string) error {
	if diff := cmp.Diff(tc.Output, output); diff != "" {
		return fmt.Errorf("output mismatch in test case %s (-expected +got):\n%s", tc.Name, diff)
	}
	return nil
}

// CompareJSONOutput compares the obtained JSON output with the expected JSON output.
func (tc *TestCase) CompareJSONOutput(output any) error {
	var expected, actual interface{}
	if err := json.Unmarshal([]byte(tc.Output), &expected); err != nil {
		return fmt.Errorf("failed to unmarshal expected JSON in test case %s: %v", tc.Name, err)
	}

	actualBytes, err := json.Marshal(output)
	if err != nil {
		return fmt.Errorf("failed to marshal actual JSON in test case %s: %v", tc.Name, err)
	}
	rtx.Must(json.Unmarshal(actualBytes, &actual))

	if diff := cmp.Diff(expected, actual); diff != "" {
		return fmt.Errorf("JSON output mismatch in test case %s (-expected +got):\n%s", tc.Name, diff)
	}
	return nil
}

// trimEmptyLines trims leading and trailing empty lines from the input string.
func trimEmptyLines(s string) string {
	lines := strings.Split(s, "\n")
	start, end := 0, len(lines)

	for start < end && strings.TrimSpace(lines[start]) == "" {
		start++
	}
	for end > start && strings.TrimSpace(lines[end-1]) == "" {
		end--
	}

	return strings.Join(lines[start:end], "\n")
}
