package device

import (
	"testing"
)

func TestInterface_Preface(t *testing.T) {
	type fields struct {
		Name string
		IP   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should return Name",
			fields: fields{
				Name: "Foo",
				IP:   "Bar",
			},
			want: "Foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				Name: tt.fields.Name,
				IP:   tt.fields.IP,
			}
			if got := i.Preface(); got != tt.want {
				t.Errorf("Interface.Preface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_Remainder(t *testing.T) {
	type fields struct {
		Name string
		IP   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should return Name",
			fields: fields{
				Name: "Foo",
				IP:   "Bar",
			},
			want: "Bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				Name: tt.fields.Name,
				IP:   tt.fields.IP,
			}
			if got := i.Remainder(); got != tt.want {
				t.Errorf("Interface.Remainder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceIter_HasNext(t *testing.T) {
	type fields struct {
		ifaces  []*Interface
		current int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "True if element left",
			fields: fields{
				ifaces: []*Interface{
					{
						Name: "Foo",
						IP:   "127.0.0.1",
					},
				},
				current: 0,
			},
			want: true,
		},
		{
			name: "False if last element",
			fields: fields{
				ifaces: []*Interface{
					{
						Name: "Foo",
						IP:   "127.0.0.1",
					},
				},
				current: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InterfaceIter{
				ifaces:  tt.fields.ifaces,
				current: tt.fields.current,
			}
			if got := i.HasNext(); got != tt.want {
				t.Errorf("InterfaceIter.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
