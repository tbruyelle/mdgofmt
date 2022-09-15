package mdgofmt

import (
	"bytes"
	"fmt"
	"go/format"
)

var (
	snipStart = []byte("```go\n")
	snipEnd   = []byte("\n```")
)

func Format(md []byte) ([]byte, error) {
	var (
		out   []byte
		start = bytes.Index(md, snipStart) + len(snipStart)
		end   = bytes.Index(md[start:], snipEnd) + start
	)
	fmt.Printf("start %d, end %d\n", start, end)

	out = append(out, md[:start]...)
	code := md[start:end]
	fmt.Println("CODE", string(code))
	formatted, err := format.Source(code)
	if err != nil {
		return nil, err
	}
	out = append(out, formatted...)
	out = append(out, md[end:]...)
	return out, nil
}
