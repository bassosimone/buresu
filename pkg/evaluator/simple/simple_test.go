// SPDX-License-Identifier: GPL-3.0-or-later

package simple_test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/bassosimone/buresu/internal/txtartesting"
	"github.com/bassosimone/buresu/pkg/evaluator/simple"
	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/scanner"
)

func TestEval(t *testing.T) {
	testCases, err := txtartesting.LoadTestCases("testdata")
	if err != nil {
		t.Fatalf("failed to load test cases: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Scan and parse the input code
			tokens, err := scanner.Scan("input.code", bytes.NewReader([]byte(tc.Input)))
			if err != nil {
				t.Fatalf("failed to scan input code: %v", err)
			}
			nodes, err := parser.Parse(tokens)
			if err != nil {
				t.Fatalf("failed to parse input code: %v", err)
			}

			// Evaluate the parsed nodes
			ctx := context.Background()
			env := simple.NewGlobalEnvironment(os.Stdout)
			var (
				results []string
				result  visitor.Value
			)
			for _, node := range nodes {
				result, err = simple.Eval(ctx, env, node)
				if err != nil {
					t.Log("err:", err.Error())
					break
				}
				t.Log("result:", result.String())
				results = append(results, result.String())
			}

			if tc.Error != "" {
				// If an error is expected, check if the actual error matches the expected error
				if err := tc.CompareError(err); err != nil {
					t.Fatal(err)
				}
				return
			}

			// If no error is expected, check if the result matches the expected output
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			actualOutput := strings.Join(results, "\n")
			if err := tc.CompareTextOutput(actualOutput); err != nil {
				t.Fatal(err)
			}
		})
	}
}
