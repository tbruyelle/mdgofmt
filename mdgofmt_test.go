package mdgofmt_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tbruyelle/mdgofmt"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		expectedError string
	}{
		{
			name:          "unclosed snipet",
			path:          "testdata/invalidsnippet.md",
			expectedError: "unclosed snippet at character 26",
		},
		{
			name:          "unclosed snipet 2",
			path:          "testdata/invalidsnippet2.md",
			expectedError: "unclosed snippet at character 6",
		},
		{
			name:          "bad code snipet",
			path:          "testdata/invalidsnippet3.md",
			expectedError: "1:25: expected operand, found ')' (and 1 more errors)",
		},
		{
			name: "empty",
			path: "testdata/empty.md",
		},
		{
			name: "no snippets",
			path: "testdata/nosnippets.md",
		},
		{
			name: "only snippet",
			path: "testdata/onlysnippet.md",
		},
		{
			name: "2 snippets",
			path: "testdata/2snippets.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			in, err := os.ReadFile(tt.path)
			require.NoError(err)

			out, err := mdgofmt.Format(in)

			if tt.expectedError != "" {
				require.EqualError(err, tt.expectedError)
				return
			}
			require.NoError(err)
			expected, err := os.ReadFile(tt.path + ".expected")
			require.NoError(err)
			require.Equal(string(expected), string(out)) // more human readable as string
		})
	}
}
