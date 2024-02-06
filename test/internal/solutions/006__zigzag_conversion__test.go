package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestConvert(t *testing.T) {
	var tests = []struct {
		s       string
		numRows int
		want    string
	}{
		{s: "PAYPALISHIRING", numRows: 3, want: "PAHNAPLSIIGYIR"},
		{s: "PAYPALISHIRING", numRows: 4, want: "PINALSIGYAHRPI"},
		{s: "A", numRows: 1, want: "A"},
		{s: "ABCDEFGHIJKLMNOP", numRows: 2, want: "ACEGIKMOBDFHJLNP"},
		{s: "ABCDEFGHIJKLMNOP", numRows: 3, want: "AEIMBDFHJLNPCGKO"},
		{s: "ABCDEFGHIJKLMNOP", numRows: 4, want: "AGMBFHLNCEIKODJP"},
	}

	for _, test := range tests {
		got := solutions.Convert(test.s, test.numRows)
		if got != test.want {
			t.Fatalf("Convert(%s, %d) == %s, want %s", test.s, test.numRows, got, test.want)
		}
	}
}
