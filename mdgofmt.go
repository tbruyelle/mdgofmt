package mdgofmt

import (
	"bytes"
	"go/format"
)

var (
	snipStart = []byte("```go\n")
	snipEnd   = []byte("\n```")
)

func Format(md []byte) ([]byte, error) {
	var out bytes.Buffer
	for {
		// fmt.Println("MD******\n", string(md), "MDEND")
		start := bytes.Index(md, snipStart)
		if start == -1 {
			break
		}
		start += len(snipStart)
		end := bytes.Index(md[start:], snipEnd)
		if end == -1 {
			break
		}
		end += start
		// fmt.Printf("len %d start %d, end %d\n", len(md), start, end)

		out.Write(md[:start])
		code := md[start:end]
		// fmt.Println("CODE******\n", string(code), "CODEEND")
		formatted, err := format.Source(code)
		if err != nil {
			return nil, err
		}
		out.Write(formatted)
		out.Write(snipEnd)
		md = md[end+len(snipEnd):]
	}
	out.WriteByte('\n')
	return out.Bytes(), nil
}
