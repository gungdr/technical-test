package anagram

import (
	"testing"
)

func Test_anagramGroup(t *testing.T) {
	tests := []struct {
		name       string
		words      []string
		wantLength int
	}{
		{
			"should return anagram",
			[]string{
				"kita",
				"atik",
				"tika",
				"aku",
				"kia",
				"makan",
				"kua",
			},
			4,
		},
		{
			"should return anagram",
			[]string{
				"ad",
				"bc",
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagramGroup(tt.words); len(got) != tt.wantLength {
				t.Errorf("anagramGroup() = %v, want %v", got, tt.wantLength)
			}
		})
	}
}
