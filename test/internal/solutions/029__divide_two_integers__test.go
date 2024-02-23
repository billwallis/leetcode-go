package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestDivide(t *testing.T) {
	var tests = []struct {
		dividend int
		divisor  int
		want     int
	}{
		{dividend: -1, divisor: 1, want: -1},
		{dividend: 0, divisor: 2, want: 0},
		{dividend: 0, divisor: 3, want: 0},
		{dividend: 3, divisor: 3, want: 1},
		{dividend: 10, divisor: 3, want: 3},
		{dividend: 7, divisor: -3, want: -2},
		{dividend: -7, divisor: -3, want: 2},
		{dividend: -2147483648, divisor: -1, want: 2147483647},
		{dividend: -2147483648, divisor: 4, want: -536870912},
	}

	for _, test := range tests {
		got := solutions.Divide(test.dividend, test.divisor)
		if got != test.want {
			t.Fatalf("Divide(%d, %d) == %d, want %d", test.dividend, test.divisor, got, test.want)
		}
	}
}
