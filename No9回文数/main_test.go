package main

import (
	"testing"
)

func TestIsPalinddrome(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "1",
			n:    -123,
			want: false,
		},
		{
			name: "2",
			n:    0,
			want: true},

		{
			name: "3",
			n:    5,
			want: true,
		},
		{
			name: "4",
			n:    123,
			want: false,
		},
		{
			name: "5",
			n:    123321,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalinddrome(tt.n); got != tt.want {
				t.Errorf("WANR....isPalinddrome(%v) return %v,want %v", tt.n, got, tt.want)
			} else {
				t.Logf("OK..isPalinddrome(%v) return %v,want %v\n", tt.n, got, tt.want)
			}
		})
	}
}
