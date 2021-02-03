package util

import (
	"bytes"
	"fmt"
)

// Prettifyer is required to make things nicely printable
type Prettifyer interface {
	Preface() string
	Remainder() string
}

// PrettiferIter iterates over a list of Prettifyer
type PrettiferIter interface {
	Next() Prettifyer
	HasNext() bool
	Reset()
}

// MakePretty create nice output
func MakePretty(p PrettiferIter) string {
	longest := 0
	for {
		if !p.HasNext() {
			break
		}
		i := p.Next()
		if len(i.Preface()) > longest {
			longest = len(i.Preface())
		}
	}
	p.Reset()
	fmtStr := fmt.Sprintf("%%%ds %%s\n", longest)
	res := bytes.NewBuffer(nil)
	for {
		if !p.HasNext() {
			break
		}
		i := p.Next()
		res.WriteString(fmt.Sprintf(fmtStr, i.Preface(), i.Remainder()))
	}
	return res.String()
}
