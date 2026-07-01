package parser

import (
	"slices"
	"testing"
)

func TestRangeByByte(t *testing.T) {
	tests := []struct {
		name string
		src  []byte
		sep  byte
		want [][]byte
	}{
		{
			name: "base",
			src:  []byte("a,b,c"),
			sep:  ',',
			want: [][]byte{
				[]byte("a"),
				[]byte("b"),
				[]byte("c"),
			},
		},
		{
			name: "empty",
			src:  []byte(""),
			sep:  ',',
			want: [][]byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got [][]byte
			rangeByByte(tt.src, tt.sep, func(start, end int) {
				got = append(got, tt.src[start:end])
			})
			for i, want := range tt.want {
				if !slices.Equal(got[i], want) {
					t.Errorf("%s: got = %v, want %v", tt.name, got, tt.want)
				}
			}
		})
	}
}

func BenchmarkRangeByByte(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		rangeByByte([]byte("a,b,c"), ',', func(start, end int) {})
	}
}

func TestTrimSpaceBytes(t *testing.T) {
	tests := []struct {
		name string
		src  []byte
		want []byte
	}{
		{
			name: "base",
			src:  []byte(" a, b, c "),
			want: []byte("a, b, c"),
		},
		{
			name: "empty",
			src:  []byte(""),
			want: []byte(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte
			trimSpaceBytes(&tt.src)
			got = append(got, tt.src...)

			if !slices.Equal(got, tt.want) {
				t.Errorf("%s: got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkTrimSpaceBytes(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		var buf []byte
		trimSpaceBytes(&buf)
	}
}
