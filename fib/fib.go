package fib

// fib 递归求解求第n个斐波那契数列的值
func fib(n int) int {
	//base case
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// fibloop 循环求解第n个斐波那契数列的值
func fibloop(n int) int {
	//base case
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}
	return b
}
