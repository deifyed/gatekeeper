package config

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestListGetter(t *testing.T) {
	testCases := []struct {
		name string

		with         string
		expectLength int
		expectList   []string
	}{
		{
			name: "Should return expected results with one item",

			with:         "http://example.com",
			expectLength: 1,
			expectList:   []string{"http://example.com"},
		},
		{
			name: "Should return expected results with one item terminated with delimiter",

			with:         "http://example.com;",
			expectLength: 1,
			expectList:   []string{"http://example.com"},
		},
		{
			name: "Should return expected results with two items",

			with:         "http://example.com;http://test.com",
			expectLength: 2,
			expectList:   []string{"http://example.com", "http://test.com"},
		},
		{
			name: "Should return expected results with two items terminated with delimiter",

			with:         "http://example.com;http://test.com;",
			expectLength: 2,
			expectList:   []string{"http://example.com", "http://test.com"},
		},
		{
			name: "Should return expected results starting line with newline",

			with:         "\nhttp://example.com",
			expectLength: 1,
			expectList:   []string{"http://example.com"},
		},
		{
			name: "Should return expected results with two items seperated by newline",

			with:         "http://example.com;\nhttp://test.com",
			expectLength: 2,
			expectList:   []string{"http://example.com", "http://test.com"},
		},
		{
			name: "Should remove spaces",

			with:         " http://example.com; http://test.com",
			expectLength: 2,
			expectList:   []string{"http://example.com", "http://test.com"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			getter := generateListGetter(func(_ string) string {
				return tc.with
			}, ";")

			result := getter("", []string{})

			assert.Equal(t, tc.expectLength, len(result))
			assert.Equal(t, tc.expectList, result)
		})
	}
}
