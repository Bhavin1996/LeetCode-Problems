package leetcode

func gcd(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	}
	return gcd(num2, num1%num2)
}
