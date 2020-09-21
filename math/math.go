package math

// Sum 对数字切片求和
func Sum(numbers[]int) int {
	sum := 0
	// BUG！！！
	for n := range numbers {
		sum += n
	}
	return sum
}