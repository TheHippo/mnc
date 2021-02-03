package transfer

import (
	"hash"
	"io"
	"testing"
)

func TestMD5Reader_Hashing(t *testing.T) {
	type fields struct {
		input   io.Reader
		hashing bool
		md5     hash.Hash
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Should be true",
			fields: fields{
				input:   nil,
				hashing: true,
				md5:     nil,
			},
			want: true,
		},
		{
			name: "Should be true",
			fields: fields{
				input:   nil,
				hashing: false,
				md5:     nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &MD5Reader{
				input:   tt.fields.input,
				hashing: tt.fields.hashing,
				md5:     tt.fields.md5,
			}
			if got := i.Hashing(); got != tt.want {
				t.Errorf("MD5Reader.Hashing() = %v, want %v", got, tt.want)
			}
		})
	}
}
