package util

import (
	"testing"
)

type element struct {
	a string
	b string
}

func (e *element) Preface() string {
	return e.a
}

func (e *element) Remainder() string {
	return e.b
}

type iter struct {
	current  int
	elements []*element
}

func (i *iter) HasNext() bool {
	return i.current < len(i.elements)
}

func (i *iter) Next() Prettifyer {
	i.current++
	return i.elements[i.current-1]
}

func (i *iter) Reset() {
	i.current = 0
}

func TestMakePretty(t *testing.T) {
	type args struct {
		p PrettiferIter
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple test",
			args: args{
				p: &iter{
					current: 0,
					elements: []*element{
						{
							a: "x",
							b: "x",
						},
						{
							a: "xx",
							b: "x",
						},
					},
				},
			},
			want: ` x x
xx x
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakePretty(tt.args.p); got != tt.want {
				t.Errorf("MakePretty() = %v, want %v", got, tt.want)
			}
		})
	}
}
