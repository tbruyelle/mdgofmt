package mdgofmt

import (
	"fmt"
	"go/format"
	"strings"
)

var (
	snipStart = "```go\n"
	snipEnd   = "\n```"
)

func Format(md string) (string, error) {
	var (
		out   strings.Builder
		start = strings.Index(md, snipStart) + len(snipStart)
		end   = strings.Index(md[start:], snipEnd) + start
	)
	fmt.Printf("start %d, end %d\n", start, end)

	out.WriteString(md[:start])
	// fmt.Println(out)
	code := md[start:end]
	fmt.Println("CODE", code)
	formatted, err := format.Source([]byte(code))
	if err != nil {
		return "", err
	}
	out.Write(formatted)
	out.WriteString(md[end:])
	return out.String(), nil
}
