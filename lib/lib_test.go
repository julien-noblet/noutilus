package lib

import "testing"

func TestRandIntMax(t *testing.T) {
	tests := []struct {
		name string
		max  int
		err  bool
	}{
		// {name: "max=0", max: 0, err: true}, // not possible it panic!
		{name: "max=1", max: 1, err: false},
		{name: "max=123456", max: 123456, err: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := RandIntMax(tt.max); got > tt.max || (err == nil && tt.err) {
				t.Errorf("RandIntMax(%v) = %v,%v want %v>=%v and err=%v", tt.max, got, err, tt.max, got, tt.err)
			}
		})
	}
}

func TestCheckLetters(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		limit  int
		unique bool
		want   bool
	}{
		// TODO: Add test cases.
		{name: "toto not unique limit=10", s: "toto", limit: 10, unique: false, want: true},
		{name: "toto unique limit=10", s: "toto", limit: 10, unique: true, want: true},
		{name: "tototototo not unique limit=10", s: "tototototo", limit: 10, unique: false, want: true},
		{name: "tototototo unique limit=10", s: "tototototo", limit: 10, unique: true, want: true},
		{name: "totototototo not unique limit=10", s: "totototototo", limit: 10, unique: false, want: false},
		{name: "totototototo unique limit=10", s: "totototototo", limit: 10, unique: true, want: true},
		{name: "azerty not unique limit=10", s: "azerty", limit: 10, unique: false, want: true},
		{name: "azerty unique limit=10", s: "azerty", limit: 10, unique: true, want: true},
		{name: "azertyuiop not unique limit=10", s: "azertyuiop", limit: 10, unique: false, want: true},
		{name: "azertyuiop unique limit=10", s: "azertyuiop", limit: 10, unique: true, want: true},
		{name: "azertyuiopq not unique limit=10", s: "azertyuiopq", limit: 10, unique: false, want: false},
		{name: "azertyuiopq unique limit=10", s: "azertyuiopq", limit: 10, unique: true, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckLetters(tt.s, tt.limit, tt.unique); got != tt.want {
				t.Errorf("CheckLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
