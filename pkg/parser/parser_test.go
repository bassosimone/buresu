// SPDX-License-Identifier: GPL-3.0-or-later

package parser_test

import (
	"strings"
	"testing"

	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/scanner"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
		shouldFail     bool
		expectedError  string
	}{
		// block tests
		{
			input:          "(block true false)",
			expectedOutput: "(block true false)",
			shouldFail:     false,
		},
		{
			input:          "(block",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:6: parser: unexpected token EOF",
		},
		{
			input:          "(block)",
			expectedOutput: "()",
			shouldFail:     false,
		},
		{
			input:          "(lambda () (block true (return! false) false))",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:12: parser: unreachable code",
		},

		// call tests
		{
			input:          "(sum 1 2 3)",
			expectedOutput: "(sum 1 2 3)",
			shouldFail:     false,
		},
		{
			input:          "(sum 1 2 3))",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:12: parser: unexpected token CLOSE",
		},
		{
			input:          "(sum",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:4: parser: unexpected token EOF",
		},

		// cond tests
		{
			input:          "(cond (true \"It's true!\") (false \"It's false!\") (else \"Neither true nor false!\"))",
			expectedOutput: "(cond (true \"It's true!\") (false \"It's false!\") (else \"Neither true nor false!\"))",
			shouldFail:     false,
		},
		{
			input:          "(cond (true 1) (else",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:20: parser: unexpected token EOF",
		},
		{
			input:          "(cond (true 1) (else 0",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:22: parser: expected token CLOSE, found EOF",
		},
		{
			input:          "(cond (",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:7: parser: unexpected token EOF",
		},
		{
			input:          "(cond (true",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:11: parser: unexpected token EOF",
		},
		{
			input:          "(cond (true 1",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:13: parser: expected token CLOSE, found EOF",
		},
		{
			input:          "(cond (true 1)",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:14: parser: expected token CLOSE, found EOF",
		},
		{
			input:          "(cond)",
			expectedOutput: "()",
			shouldFail:     false,
		},
		{
			input:          "(cond (else true))",
			expectedOutput: "true",
			shouldFail:     false,
		},
		{
			input:          "(cond (true 1))",
			expectedOutput: "(cond (true 1) (else ()))",
			shouldFail:     false,
		},

		// define & set tests
		{
			input:          "(set! x 42)",
			expectedOutput: "(set! x 42)",
			shouldFail:     false,
		},
		{
			input:          "(define x 42)",
			expectedOutput: "(define x 42)",
			shouldFail:     false,
		},
		{
			input:          "(define",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:7: parser: expected token ATOM, found EOF",
		},
		{
			input:          "(define x",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:9: parser: unexpected token EOF",
		},
		{
			input:          "(define x 42",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:12: parser: expected token CLOSE, found EOF",
		},

		// if tests
		{
			input:          "(if true 1 0)",
			expectedOutput: "(cond (true 1) (else 0))",
			shouldFail:     false,
		},
		{
			input:          "(if true 1)",
			expectedOutput: "(cond (true 1) (else ()))",
			shouldFail:     false,
		},
		{
			input:          "(if",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:3: parser: unexpected token EOF",
		},
		{
			input:          "(if true",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:8: parser: unexpected token EOF",
		},
		{
			input:          "(if true",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:8: parser: unexpected token EOF",
		},
		{
			input:          "(if true else",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:13: parser: unexpected token EOF",
		},
		{
			input:          "(if true else (",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:15: parser: unexpected token EOF",
		},

		// lambda tests
		{
			input:          "(lambda x)",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:9: parser: expected token OPEN, found ATOM",
		},
		{
			input:          "(lambda (x) \"This is a lambda function\" 42)",
			expectedOutput: "(lambda (x) \"This is a lambda function\" 42)",
			shouldFail:     false,
		},
		{
			input:          "(lambda",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:7: parser: expected token OPEN, found EOF",
		},
		{
			input:          "(lambda (",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:9: parser: unexpected token EOF",
		},
		{
			input:          "(lambda (()",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:1: parser: lambda parameter name must be a symbol",
		},
		{
			input:          "(lambda (x))",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:12: parser: unexpected token CLOSE",
		},
		{
			input:          "(lambda (x) x",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:13: parser: expected token CLOSE, found EOF",
		},
		{
			input:          "(lambda (x x) x)",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:1: parser: lambda parameter \"x\" is duplicated",
		},

		// quote tests
		{
			input:          "(quote (1 2 3))",
			expectedOutput: "(quote (1 2 3))",
			shouldFail:     false,
		},
		{
			input:          "(quote",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:6: parser: unexpected token EOF",
		},
		{
			input:          "(quote (1 2 3)",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:14: parser: expected token CLOSE, found EOF",
		},

		// return tests
		{
			input:          "(return! 42)",
			expectedOutput: "(return! 42)",
			shouldFail:     true,
			expectedError:  "<stdin>:1:1: parser: return! outside of lambda",
		},
		{
			input:          "(block (return! 42))",
			expectedOutput: "(return! 42)",
			shouldFail:     true,
			expectedError:  "<stdin>:1:8: parser: return! outside of lambda",
		},
		{
			input:          "(lambda () (return! 42))",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:12: parser: statement not allowed in this context",
		},
		{
			input:          "(lambda () (lambda () (return! 42)))",
			expectedOutput: "(lambda () \"\" (lambda () \"\" (return! 42)))",
			shouldFail:     true,
			expectedError:  "<stdin>:1:23: parser: statement not allowed in this context",
		},
		{
			input:          "(lambda () (return!",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:19: parser: unexpected token EOF",
		},
		{
			input:          "(lambda () (return! 42",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:22: parser: expected token CLOSE, found EOF",
		},

		// while tests
		{
			input:          "(while true 42)",
			expectedOutput: "(while true 42)",
			shouldFail:     false,
		},
		{
			input:          "(while true 42.42)",
			expectedOutput: "(while true 42.42)",
			shouldFail:     false,
		},
		{
			input:          "(while true ())",
			expectedOutput: "(while true ())",
			shouldFail:     false,
		},
		{
			input:          "(while",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:6: parser: unexpected token EOF",
		},
		{
			input:          "(while true",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:11: parser: unexpected token EOF",
		},
		{
			input:          "(while true ()",
			expectedOutput: "",
			shouldFail:     true,
			expectedError:  "<stdin>:1:14: parser: expected token CLOSE, found EOF",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			tokens, err := scanner.Scan("<stdin>", strings.NewReader(test.input))
			if err != nil {
				t.Fatalf("unexpected error scanning input: %v", err)
			}

			nodes, err := parser.Parse(tokens)
			if test.shouldFail {
				if err == nil {
					t.Errorf("expected error but got none")
				} else if err.Error() != test.expectedError {
					t.Errorf("expected error %q but got %q", test.expectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error parsing input: %v", err)
			}

			var output []string
			for _, node := range nodes {
				output = append(output, node.String())
			}

			result := strings.Join(output, " ")
			if result != test.expectedOutput {
				t.Errorf("expected output %q but got %q", test.expectedOutput, result)
			}
		})
	}
}
