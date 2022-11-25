package mdgofmt

import (
	"bytes"
	"fmt"
	"go/format"
)

var (
	snipStart = []byte("```go")
	snipEnd   = []byte("```")

	replacements = [][][]byte{
		{[]byte("{ModulePath}"), []byte("ModulePath")},
		{[]byte("{BinaryNamePrefix}"), []byte("BinaryNamePrefix")},
	}
)

func Format(md []byte) ([]byte, error) {
	var (
		out       bytes.Buffer
		fileIndex int
	)
	for {
		start := bytes.Index(md, snipStart)
		if start == -1 {
			out.Write(md)
			break
		}
		// move start until it reach the end of snipStart line
		// (we may need to skip snippet's attributes)
		for start += len(snipStart); md[start] != '\n'; start++ {
		}
		start++ // skip final \n

		end := bytes.Index(md[start:], snipEnd)
		if end == -1 {
			return nil, fmt.Errorf("unclosed snippet at character %d", start+fileIndex)
		}
		end += start
		// fmt.Printf("len %d start %d, end %d\n", len(md), start, end)

		out.Write(md[:start])
		code := md[start:end]
		for i := range replacements {
			code = bytes.ReplaceAll(code, replacements[i][0], replacements[i][1])
		}
		// fmt.Println("CODE******\n", string(code), "CODEEND")
		formatted, err := format.Source(code)
		if err != nil {
			return nil, fmt.Errorf("format source at %d: %w\n%s", start+fileIndex, err, code)
		}
		for i := range replacements {
			formatted = bytes.ReplaceAll(formatted, replacements[i][1], replacements[i][0])
		}
		out.Write(formatted)
		out.Write(snipEnd)
		// truncate md to remaining data
		skip := end + len(snipEnd)
		md = md[skip:]
		fileIndex += skip
	}
	return out.Bytes(), nil
}
