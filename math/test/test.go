package main

import (
	"fmt"

	"github.com/longyue0521/go-testing-notes/math"
)

// TestSum 用于测试math包的Sum方法
func TestSum() {
	// 测试用例1
	sum := math.Sum([]int{0, 1, 2})
	if sum != 3 {
		// 个性化提示信息，以便快速找到测试用例
		msg := fmt.Sprintf("FAIL 1: Wanted 3 but Got %d", sum)
		panic(msg)
	}
	fmt.Println("PASS 1")

	// 测试用例2
	sum = math.Sum([]int{10, -1, 2})
	if sum != 11 {
		// 个性化提示信息，以便快速找到测试用例
		msg := fmt.Sprintf("FAIL 2: Wanted 11 but Got %d", sum)
		panic(msg)
	}
	fmt.Println("PASS 2")
}

func main() {
	TestSum()
}
