// SPDX-License-Identifier: GPL-3.0-or-later

package txtartesting_test

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/bassosimone/buresu/internal/txtartesting"
)

func TestLoadTestCases_Successful(t *testing.T) {
	testCases, err := txtartesting.LoadTestCases("testdata")
	if err != nil {
		t.Fatal(err)
	}
	if len(testCases) < 1 {
		t.Fatal("expected at least one test case, got none")
	}
}

func TestLoadTestCases_NonexistentDir(t *testing.T) {
	testCases, err := txtartesting.LoadTestCases("nonexistent")
	if err == nil {
		t.Fatalf("expected error for nonexistent directory, but got no error")
	}
	if len(testCases) != 0 {
		t.Fatal("expected no test cases, got", len(testCases))
	}
}

func TestLoadTestCases_DirWithErrors(t *testing.T) {
	testCases, err := txtartesting.LoadTestCases(filepath.Join("testdata", "failurecases"))
	if err == nil {
		t.Fatalf("expected error for malformed test suite, but got no error")
	}
	if len(testCases) != 0 {
		t.Fatal("expected no test cases, got", len(testCases))
	}
}

func TestLoadSingleTestCase_ValidTest(t *testing.T) {
	tc, err := txtartesting.LoadSingleTestCase(filepath.Join(
		"testdata", "valid-test.txtar"))
	if err != nil {
		t.Fatal(err)
	}

	if tc.Name != "valid-test.txtar" {
		t.Errorf("expected test case name to be 'valid-test.txtar', got '%s'", tc.Name)
	}

	expectedInput := "(def x 42)"
	if tc.Input != expectedInput {
		t.Errorf("expected input to be '%s', got '%s'", expectedInput, tc.Input)
	}

	expectedOutput := `{
  "type": "DefineExpr",
  "symbol": "x",
  "expr": {
    "type": "IntLiteral",
    "value": "42"
  }
}`
	if tc.Output != expectedOutput {
		t.Errorf("expected output to be '%s', got '%s'", expectedOutput, tc.Output)
	}

	if tc.Error != "" {
		t.Errorf("expected error to be empty, got '%s'", tc.Error)
	}
}

func TestLoadSingleTestCase_FailureCases_MissingOutputAndError(t *testing.T) {
	_, err := txtartesting.LoadSingleTestCase(filepath.Join(
		"testdata", "failurecases", "missing-output-and-error.txtar"))
	if err == nil {
		t.Fatalf("expected error for missing output and error, but got no error")
	}
}

func TestLoadSingleCase_FailureCases_UnexpectedSection(t *testing.T) {
	_, err := txtartesting.LoadSingleTestCase(filepath.Join(
		"testdata", "failurecases", "unexpected-section.txtar"))
	if err == nil {
		t.Fatalf("expected error for unexpected section, but got no error")
	}
}

func TestLoadSingleCase_FailureCases_NonexistentFile(t *testing.T) {
	_, err := txtartesting.LoadSingleTestCase(filepath.Join(
		"testdata", "failurecases", "nonexistent-file.txtar"))
	if err == nil {
		t.Fatalf("expected error for nonexistent file, but got no error")
	}
}

func TestCompareError(t *testing.T) {
	tc := &txtartesting.TestCase{
		Name:  "test",
		Error: "expected error",
	}

	t.Run("Success", func(t *testing.T) {
		if err := tc.CompareError(errors.New("expected error")); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		if err := tc.CompareError(errors.New("unexpected error")); err == nil {
			t.Fatal("expected error, got none")
		}
	})

	t.Run("ExpectedErrorButGotNone", func(t *testing.T) {
		if err := tc.CompareError(nil); err == nil {
			t.Fatal("expected error, got none")
		}
	})

	t.Run("NoErrorExpectedButGotOne", func(t *testing.T) {
		tc := &txtartesting.TestCase{
			Name:  "test",
			Error: "",
		}
		if err := tc.CompareError(errors.New("unexpected error")); err == nil {
			t.Fatal("expected no error, but got one")
		}
	})

	t.Run("NoErrorExpectedAndGotNone", func(t *testing.T) {
		tc := &txtartesting.TestCase{
			Name:  "test",
			Error: "",
		}
		if err := tc.CompareError(nil); err != nil {
			t.Fatal("expected no error, but got one")
		}
	})
}

func TestCompareTextOutput(t *testing.T) {
	tc := &txtartesting.TestCase{
		Name:   "test",
		Output: "expected output",
	}

	t.Run("Success", func(t *testing.T) {
		if err := tc.CompareTextOutput("expected output"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		if err := tc.CompareTextOutput("unexpected output"); err == nil {
			t.Fatal("expected error, got none")
		}
	})
}

func TestCompareJSONOutput(t *testing.T) {
	tc := &txtartesting.TestCase{
		Name:   "test",
		Output: `{"key": "value"}`,
	}

	t.Run("Success", func(t *testing.T) {
		if err := tc.CompareJSONOutput(map[string]string{"key": "value"}); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		if err := tc.CompareJSONOutput(map[string]string{"key": "unexpected value"}); err == nil {
			t.Fatal("expected error, got none")
		}
	})

	t.Run("InvalidExpectedJSON", func(t *testing.T) {
		tc := &txtartesting.TestCase{
			Name:   "test",
			Output: `invalid json`,
		}
		if err := tc.CompareJSONOutput(map[string]string{"key": "value"}); err == nil {
			t.Fatal("expected error for invalid expected JSON, got none")
		}
	})

	t.Run("UnmarshalableOutput", func(t *testing.T) {
		if err := tc.CompareJSONOutput(make(chan any)); err == nil {
			t.Fatal("expected error for unmarshalable output, got none")
		}
	})
}
