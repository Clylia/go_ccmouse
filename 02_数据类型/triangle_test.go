package main

/**
|-------------------------------------------------------------------------------------------------
|
|	测试，表格驱动测试
|
|	命令行中使用 go test . 执行当前目录下的所有测试
|
|-------------------------------------------------------------------------------------------------
*/

import (
	"testing"
)

func TestTriangle(t *testing.T)  {
	tests := []struct {a, b, c int} {
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
