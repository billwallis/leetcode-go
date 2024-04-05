package solutions_test

import (
	"github.com/Bilbottom/leetcode-go/internal/solutions"
	"testing"
)

func TestMultiple(t *testing.T) {
	t.Skip("Not correctly implemented yet...")

	var tests = []struct {
		num1 string
		num2 string
		want string
	}{
		{num1: "2", num2: "3", want: "6"},
		{num1: "123", num2: "456", want: "56088"},
		{num1: "498828660196", num2: "840477629533", want: "419254329864656431168468"},
	}

	for _, test := range tests {
		got := solutions.Multiply(test.num1, test.num2)
		if got != test.want {
			t.Fatalf("Multiply(%s, %s) == %s, want %s", test.num1, test.num2, got, test.want)
		}
	}
}
