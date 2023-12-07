package main

func main() {
	romanToInt("LLX")
}

var digits = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			sum += digits[s[i]]
			break
		}
		if digits[s[i]] < digits[s[i+1]] {
			sum -= digits[s[i]]
		} else {
			sum += digits[s[i]]
		}
	}
	return sum
}
