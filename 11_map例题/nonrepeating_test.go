package main


/**
|-------------------------------------------------------------------------------------------------
|
|	测试，表格驱动测试
|
|	普通测试，性能测试，性能调优（pprof）
|
|-------------------------------------------------------------------------------------------------
*/

import (
	"testing"
)

// 普通测试
func TestSubstr(t *testing.T)  {
	tests := []struct{
		s string
		ans int
	} {
		// Normal cases
		{"lalala", 2},
		{"abcddcac", 4},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbb", 1},
		{"abcdef", 6},

		// chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		if actual := getNonRepeatStrByRune(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual,tt.s, tt.ans)
		}
	}
}

// 性能测试
// 性能测试需要多次执行，具体执行多少次，是需要一个算法的，这个算法有 b.N 来得出，只需要在代码中实现他即可
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	// 将测试数据加长，添加测试强度
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 8

	// 查看字符串长度
	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()	// 重置时间，以免上面的代码影响测试结果有误差

	// 执行次数由 b.N 得出
	for i := 0; i < b.N; i++ {
		if actual := getNonRepeatStrByRune(s); actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
