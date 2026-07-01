package parser

import (
	"net/url"
	"testing"
)

func TestAddTitle(t *testing.T) {
	tests := []struct {
		name string
		find []byte
		want string
	}{
		{
			name: "base",
			find: []byte("a,b,c"),
			want: "NAME:(a OR b OR c)",
		},
		{
			name: "empty",
			find: []byte(""),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ""
			AddTitle(&got, tt.find)

			if got != tt.want {
				t.Errorf("%s: got = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestAddSalary(t *testing.T) {
	tests := []struct {
		name string
		sal  []byte
		want string
	}{
		{
			name: "base",
			sal:  []byte(" 1000 "),
			want: "1000",
		},
		{
			name: "empty",
			sal:  []byte(""),
			want: "",
		},
	}

	got := ""
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddSalary(&got, tt.sal)

			if got != tt.want {
				t.Errorf("%s: got = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkAddSalary(b *testing.B) {
	got := ""
	slice := []byte(" 1000 ")
	b.ResetTimer()
	for b.Loop() {
		AddSalary(&got, slice)
	}
}

func TestAddExp(t *testing.T) {
	tests := []struct {
		name string
		exp  []byte
		want string
	}{
		{
			name: "base",
			exp:  []byte("noExp,between1And3"),
			want: "experience=noExp&experience=between1And3",
		},
		{
			name: "empty",
			exp:  []byte(""),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := url.Values{}
			AddExp(&got, tt.exp)

			if got.Encode() != tt.want {
				t.Errorf("%s: got = %q, want %q", tt.name, got.Encode(), tt.want)
			}
		})
	}
}

func BenchmarkAddExp(b *testing.B) {
	got := url.Values{}
	slice := []byte("noExp,between1And3")
	b.ResetTimer()
	for b.Loop() {
		AddExp(&got, slice)
	}
}

func TestAddSchedule(t *testing.T) {
	tests := []struct {
		name string
		sch  []byte
		want string
	}{
		{
			name: "base",
			sch:  []byte("a,b,c"),
			want: "a,b,c",
		},
		{
			name: "empty",
			sch:  []byte(""),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ""
			AddSchedule(&got, tt.sch)

			if got != tt.want {
				t.Errorf("%s: got = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkAddSchedule(b *testing.B) {
	got := ""
	slice := []byte("a,b,c")
	b.ResetTimer()
	for b.Loop() {
		AddSchedule(&got, slice)
	}
}
