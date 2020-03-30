package main

func main() {
	println(reverse(1534236469))
}

func reverse(x int) int {
	var a int
	if x == 0 {
		return 0
	}
	for {
		a = a*10 + x%10
		x = x / 10
		if x < 1 && a > 0 || -x < 1 && -a > 0 {
			break
		}
	}
	xt := int32(a)
	if int(xt) < a && a > 0 || int(xt) > a {
		return 0
	}
	return a
}
