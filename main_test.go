package main

import (
	"testing"
)

func TestAddFindTitle(t *testing.T) {
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
			want: "NAME:()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ""
			addFindTitle(&got, tt.find)

			if got != tt.want {
				t.Errorf("%s: got = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
