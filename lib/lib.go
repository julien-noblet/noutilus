package lib

import (
	"crypto/rand"
	"errors"
	"math/big"
	mrand "math/rand"
	"regexp"
	"sort"
	"strings"
)

var ErrRandomFail = errors.New(`random fail`)

var ErrPrimaryNumber = errors.New("primary number?")

func ReduceUniqueLetters(s string) string {
	ret := ""
	mapS := make(map[rune]bool)

	for _, rr := range s {
		mapS[rr] = true
	}

	for k := range mapS {
		ret += string(k)
	}

	return ret
}

func CountUniqueLetters(s string) int {
	mapS := make(map[rune]bool)
	for _, rr := range s {
		mapS[rr] = true
	}

	return len(mapS)
}

func CheckLetters(s string, limit int, unique bool) bool {
	if unique {
		return (CountUniqueLetters(s) <= limit)
	}

	return (len(s) <= limit)
}

func AddNoise(s string, size int) string {
	noise, _ := ShuffleWord(`abcdefghijklmnopqrstuvwxyz`)
	ret := s

	for len(ret) < size {
		ret = ReduceUniqueLetters(ret + string(noise[len(noise)-1]))
		noise = noise[:len(noise)-1]
	}

	return ret
}

func SplitText(s string) []string {
	re := regexp.MustCompile(`[A-Za-z1-9']+|[':?().,!\\ ]`)

	return re.FindAllString(s, -1)
}

func ShuffleWord(s string) (string, error) {
	word := strings.Split(s, "")

	mrand.Shuffle(len(word), func(i, j int) {
		word[i], word[j] = word[j], word[i]
	})

	return strings.Join(word, ""), nil
}

func Numerize(s, ttf string) []int {
	ret := []int{}
	for _, v := range ttf {
		ret = append(ret, strings.IndexRune(s, v))
	}

	return ret
}

func ConcatInt(s []int) int {
	res := 0
	op := 1

	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}

	return res
}

func RandIntMin(min int) (int, error) {
	r, e := rand.Int(rand.Reader, big.NewInt(int64(min)))
	if e != nil {
		return 0, e
	}

	if r.IsInt64() {
		return int(r.Int64()) + min, e
	}

	return 0, ErrRandomFail
}

func RandIntMax(max int) (int, error) {
	r, e := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if e != nil {
		return 0, e
	}

	if r.IsInt64() {
		return int(r.Int64()), e
	}

	return 0, ErrRandomFail
}

func DecompositionNombresPremiers(n int) []int {
	// TODO : Translate in english!!!
	if n < 2 { //nolint:gomnd // it doesn't work if n<2 :p
		return []int{}
	}

	f := []int{}

	for i := 2; i <= n; i++ { //nolint:gomnd // start at 2 :/
		for n%i == 0 {
			f = append(f, i)
			n /= i
		}
	}

	return f
}

func MultIntSlice(ns []int) int {
	ret := 1
	for _, v := range ns {
		ret *= v
	}

	return ret
}

func Find2Factors(n int) (first, second int, err error) {
	factors := DecompositionNombresPremiers(n)
	if len(factors) < 2 { //nolint:gomnd // If len(factors) < 2 => primary number :p
		return 0, 0, ErrPrimaryNumber
	}

	sort.Ints(factors)
	first = factors[len(factors)-1]
	second = MultIntSlice(factors[:len(factors)-1])

	return first, second, nil
}
