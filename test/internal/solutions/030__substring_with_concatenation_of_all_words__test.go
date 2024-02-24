package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	var tests = []struct {
		s     string
		words []string
		want  []int
	}{
		{s: "abcd", words: []string{"b", "c", "d"}, want: []int{1}},
		{s: "barfoothefoobarman", words: []string{"foo", "bar"}, want: []int{0, 9}},
		{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "word"}, want: []int{}},
		{s: "barfoofoobarthefoobarman", words: []string{"bar", "foo", "the"}, want: []int{6, 9, 12}},
	}

	for _, test := range tests {
		got := solutions.FindSubstring(test.s, test.words)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("FindSubstring(%s, %v) == %v, want %v", test.s, test.words, got, test.want)
		}
	}
}
