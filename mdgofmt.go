package mdgofmt

import (
	"bytes"
	"fmt"
	"go/format"
)

var (
	snipStart = []byte("```go\n")
	snipEnd   = []byte("```")
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
		start += len(snipStart)

		end := bytes.Index(md[start:], snipEnd)
		if end == -1 {
			return nil, fmt.Errorf("unclosed snippet at character %d", start+fileIndex)
		}
		end += start
		// fmt.Printf("len %d start %d, end %d\n", len(md), start, end)

		out.Write(md[:start])
		code := md[start:end]
		// fmt.Println("CODE******\n", string(code), "CODEEND")
		formatted, err := format.Source(code)
		if err != nil {
			return nil, fmt.Errorf("format source at %d: %w", start+fileIndex, err)
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
