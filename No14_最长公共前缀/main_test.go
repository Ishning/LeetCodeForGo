package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		desc string
		strs []string
		want string
	}{
		{
			desc: "1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			desc: "2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			desc: "3",
			strs: []string{"long", "longboard", "longevity"},
			want: "long",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if resulStr := longestCommonPrefix(tC.strs); resulStr != tC.want {
				t.Errorf("WRAN...longestCommonPrefix(%v) return %v, want %v", tC.strs, resulStr, tC.want)
			} else {
				t.Logf("OK...longestCommonPrefix(%v) return %v, want %v", tC.strs, resulStr, tC.want)
			}
		})
	}
}
