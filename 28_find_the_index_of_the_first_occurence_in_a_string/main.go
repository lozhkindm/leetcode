package main

func main() {

}

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	l := 0
	r := len(needle)
	for r <= len(haystack) {
		if haystack[l:r] == needle {
			return l
		}
		l++
		r++
	}
	return -1
}
