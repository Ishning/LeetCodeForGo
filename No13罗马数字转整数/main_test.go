package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name  string
		roman string
		want  int
	}{
		{
			name:  "1",
			roman: "III",
			want:  3,
		},
		{
			name:  "2",
			roman: "IV",
			want:  4,
		},
		{
			name:  "3",
			roman: "IX",
			want:  9,
		},
		{
			name:  "4",
			roman: "LVIII",
			want:  58, //L = 50, V= 5, III = 3
		},
		{
			name:  "5",
			roman: "MCMXCIV",
			want:  1994, //M = 1000, CM = 900, XC = 90, IV = 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if resultInt := romanToInt(tt.roman); resultInt != tt.want {
				t.Errorf("WANR....romanToInt(%v) return %v, want %v", tt.roman, resultInt, tt.want)
			} else {
				t.Logf("OK...romanToInt(%v) return %v, want %v", tt.roman, resultInt, tt.want)
			}
		})
	}
}
