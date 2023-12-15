package main

func main() {
	mySqrt(0)
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}

	l := 2
	r := x / 2

	for l < r {
		length := ((l + r) / 2) + 1
		lines := x / length
		diff := length - lines
		if -1 <= diff && diff < 1 {
			return length
		}
		if diff > 0 {
			r = length - 1
		} else {
			l = length + 1
		}
	}

	return r
}
