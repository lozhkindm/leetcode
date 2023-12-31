package main

func main() {
	isValid("(({(}")
}

func isValid(s string) bool {
	stack := make([]uint8, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != ')' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != ']' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '}' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
