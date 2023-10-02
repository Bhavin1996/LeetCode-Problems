package leetcode

import "fmt"

func findIthBit(num int, k int) int {
	mask := 1 << k
	if num&mask > 0 {
		return 1
	} else {
		return 0
	}

}

func setIthBit(num int, k int) int {
	mask := 1 << k
	num = num | mask
	return num
}

func clearIthBit(num int, k int) {
	mask := ^(1 << k)
	res := num & mask
	fmt.Println(res)
}

func numberOfSetBits(num int) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

func findNumberOfBitsToChnageA_to_B(num int, num2 int) int {
	count := 0
	res := 0
	res = num ^ num2
	count = numberOfSetBits(res)
	return count
}
