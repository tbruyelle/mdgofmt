package mdgofmt

import "go/format"

func Format(md []byte) ([]byte, error) {
	f, err := format.Source(md)
	return f, err
}
