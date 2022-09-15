package mdgofmt_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tbruyelle/mdgofmt"
)

func TestFormat(t *testing.T) {
	require := require.New(t)
	in, err := os.ReadFile("testdata/test1.md")
	require.NoError(err)

	out, err := mdgofmt.Format(in)

	require.NoError(err)
	expected, err := os.ReadFile("testdata/test1_expected.md")
	require.NoError(err)
	require.Equal(string(expected), string(out)) // more human readable as string
}
