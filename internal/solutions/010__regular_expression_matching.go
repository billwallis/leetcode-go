package solutions

/*
- `.` Matches any single character.
- `*` Matches zero or more of the preceding element.

The pattern `m*h` matches `mmmmh` because the `m` can be used 4 times:

	mmmmh
    m***h

The pattern `m*h` does not match `match` because the `m` can't be used
for the `a`:

	match
    m-  h

The pattern `m.*h` matches `match` because the `.` can be used 3 times:

	match
    m...h

The pattern `m.*m` does not match `match` because the `m` at the end
doesn't match anything in the pattern:

	match
    m....m

The pattern `.*aa.*b.*aa.*` matches `aabaa` because we can take 0 match
for each `.`:

     aa b aa
    -aa-b-aa-

The pattern `.*a.*b.*a.*` matches `aabaa` because we can take 1 match
from either `.` on either side of `b`:

     aa b aa
    -a. b .a-
     .a-b-a.
*/

// isMatch is the tenth LeetCode problem:
// - https://leetcode.com/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	return s == p
}

var IsMatch = isMatch /* For testing */
