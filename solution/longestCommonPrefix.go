package solution

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	r := []rune(strs[0])
	prefix := ""

	for i := 0; i < len(r); i++ {
		char := r[i]
		for j := 1; j < len(strs); j++ {
			r_item := []rune(strs[j])
			if len(r_item)-1 < i || r_item[i] != char {
				return prefix
			}
			if j == len(strs)-1 {
				prefix += string(char)
			}
		}
	}
	return prefix
}
