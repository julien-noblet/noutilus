package lib

import (
	"errors"
	"reflect"
	"sort"
	"testing"
)

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

func TestReduceUniqueLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "ehlo"},
		{"world", "dlorw"},
		{"", ""},
		{"aabbcc", "abc"},
		{"abcabc", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ReduceUniqueLetters(tt.input)
			// Sort the result string
			resultSorted := sortString(result)
			expectedSorted := sortString(tt.expected)
			if resultSorted != expectedSorted {
				t.Errorf("ReduceUniqueLetters(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Helper function to sort a string
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func TestAddNoise(t *testing.T) {
	tests := []struct {
		input    string
		size     int
		expected string
	}{
		{"hello", 10, "helo"},
		{"world", 10, "world"},
		{"", 5, "abcde"},
		{"aabbcc", 8, "abc"},
		{"abcabc", 6, "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := AddNoise(tt.input, tt.size)
			if len(result) != tt.size {
				t.Errorf("AddNoise(%q, %d) = %q; length = %d; want length %d", tt.input, tt.size, result, len(result), tt.size)
			}
		})
	}
}

func TestSplitText(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"Hello, world!", []string{"Hello", ",", " ", "world", "!"}},
		{"Go is great.", []string{"Go", " ", "is", " ", "great", "."}},
		{"What's up?", []string{"What's", " ", "up", "?"}},
		//{"", []string{}}, // does it make sense to test this?
		{"123", []string{"123"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := SplitText(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SplitText(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNumerize(t *testing.T) {
	tests := []struct {
		s        string
		ttf      string
		expected []int
	}{
		{"abcdef", "ace", []int{0, 2, 4}},
		{"hello", "lo", []int{2, 4}},
		{"world", "wrd", []int{0, 2, 4}},
		{"", "abc", []int{-1, -1, -1}},
		{"abc", "", []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.s+"_"+tt.ttf, func(t *testing.T) {
			result := Numerize(tt.s, tt.ttf)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Numerize(%q, %q) = %v; want %v", tt.s, tt.ttf, result, tt.expected)
			}
		})
	}
}

func TestConcatInt(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 123},
		{[]int{0, 1, 2}, 12},
		{[]int{9, 8, 7}, 987},
		{[]int{0}, 0},
		{[]int{}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := ConcatInt(tt.input)
			if result != tt.expected {
				t.Errorf("ConcatInt(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRandIntMin(t *testing.T) {
	tests := []struct {
		min      int
		expected int
	}{
		{10, 10},
		{1, 1},
		{5, 5},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result, err := RandIntMin(tt.min)
			if err != nil {
				t.Errorf("RandIntMin(%d) returned error: %v", tt.min, err)
			}
			if result < tt.min {
				t.Errorf("RandIntMin(%d) = %d; want >= %d", tt.min, result, tt.min)
			}
		})
	}
}

func TestDecompositionNombresPremiers(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{1, []int{}},
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{5, []int{5}},
		{6, []int{2, 3}},
		{7, []int{7}},
		{8, []int{2, 2, 2}},
		{9, []int{3, 3}},
		{10, []int{2, 5}},
		{12, []int{2, 2, 3}},
		{15, []int{3, 5}},
		{16, []int{2, 2, 2, 2}},
		{18, []int{2, 3, 3}},
		{20, []int{2, 2, 5}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := DecompositionNombresPremiers(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DecompositionNombresPremiers(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMultIntSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 24},
		{[]int{2, 3, 5}, 30},
		{[]int{10, 10, 10}, 1000},
		{[]int{0, 1, 2}, 0},
		{[]int{1}, 1},
		{[]int{}, 1}, // Assuming the product of an empty slice is 1
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := MultIntSlice(tt.input)
			if result != tt.expected {
				t.Errorf("MultIntSlice(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFind2Factors(t *testing.T) {
	tests := []struct {
		input       int
		expected1   int
		expected2   int
		expectedErr error
	}{
		{6, 3, 2, nil},              // 6 = 2 * 3
		{28, 7, 4, nil},             // 28 = 4 * 7
		{49, 7, 7, nil},             // 49 = 7 * 7
		{2, 0, 0, ErrPrimaryNumber}, // 2 is a prime number
		{1, 0, 0, ErrPrimaryNumber}, // 1 has no prime factors
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			first, second, err := Find2Factors(tt.input)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Find2Factors(%d) error = %v; want %v", tt.input, err, tt.expectedErr)
			}
			if first != tt.expected1 || second != tt.expected2 {
				t.Errorf("Find2Factors(%d) = (%d, %d); want (%d, %d)", tt.input, first, second, tt.expected1, tt.expected2)
			}
		})
	}
}
