package main

func main() {
	longestCommonPrefix([]string{})
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(prefix) {
			prefix = prefix[:len(strs[i])]
		}

		for j := 0; j < len(prefix); j++ {
			if strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}
