package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	// 测试用例1
	sum := Sum([]int{0, 1, 2})
	if sum != 3 {
		// 个性化提示信息，以便快速找到测试用例
		t.Errorf("FAIL 1: Wanted 3 but Got %d", sum)
	}

	// 测试用例2
	sum = Sum([]int{10, -1, 2})
	if sum != 11 {
		// 个性化提示信息，以便快速找到测试用例
		t.Errorf("FAIL 2: Wanted 11 but Got %d", sum)
	}
}
