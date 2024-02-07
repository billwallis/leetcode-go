package solutions

import (
	"math"
	"strings"
)

/*
Take the string `ABCDEFGHIJKLMNOP`.

Split over 2 rows would look like:

    A C E G I K M O
    B D F H J L N P

    0 2 4 6
    1 3 5 7

Split over 3 rows would look like:

    A   E   I   M
    B D F H J L N P
    C   G   K   O

    0   4
    1 3 5 7
    2   6

...and split over 4 rows would look like:

    A     G     M
    B   F H   L N
    C E   I K   O
    D     J     P

	0     6
    1   5 7
    2 4
    3

For n >= 2, there are n + (n - 2) == 2n - 2 items per "segment".

The items on the first row (row 0) are those whose index is 0 mod n.
then, the items on row i are those whose index is i or n - i mod n.
*/

func MapIndexToRow(number int, modulo int) int {
	n, m := float64(number), float64(modulo)
	res := math.Mod(n, m)
	return int(math.Min(res, m-res))
}

// Convert is the sixth LeetCode problem:
// - https://leetcode.com/problems/zigzag-conversion/
func Convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	modulo := (2 * numRows) - 2
	rows := make([]string, numRows)

	for index, char := range s {
		row := MapIndexToRow(index, modulo)
		rows[row] += string(char)
	}

	return strings.Join(rows, "")
}
