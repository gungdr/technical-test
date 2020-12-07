package refactor

import (
	"testing"
)

func Test_findFirstStringInBracket(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{
			"should return empty string",
			"",
			"",
		},
		{
			"should return empty string",
			"agung",
			"",
		},
		{
			"should return empty string",
			"(agung",
			"",
		},
		{
			"should return empty string",
			"agung)",
			"",
		},
		{
			"should return string",
			"(agung)",
			"agung",
		},
		{
			"should return string",
			"((agung))",
			"agung",
		},
		{
			"should return string",
			"((agung)))",
			"agung",
		},
		{
			"should return string",
			"(((agung))",
			"agung",
		},
		{
			"should return string",
			"hehe(agung)",
			"agung",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstStringInBracket(tt.word); got != tt.want {
				t.Errorf("findFirstStringInBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}
