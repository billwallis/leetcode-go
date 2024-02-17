package solutions

func opposite(bracket uint8) uint8 {
	switch bracket {
	case "("[0]:
		return ")"[0]
	case ")"[0]:
		return "("[0]
	case "{"[0]:
		return "}"[0]
	case "}"[0]:
		return "{"[0]
	case "["[0]:
		return "]"[0]
	case "]"[0]:
		return "["[0]
	default:
		panic("Unknown bracket")
	}
}

// isValid is the twentieth LeetCode problem:
// - https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	openBrackets := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == "("[0] || char == "["[0] || char == "{"[0] {
			openBrackets += string(char)
		} else if length := len(openBrackets); length > 0 && opposite(char) == openBrackets[length-1] {
			openBrackets = openBrackets[:length-1]
		} else {
			return false
		}
	}
	return openBrackets == ""
}

var IsValid = isValid /* For testing */
