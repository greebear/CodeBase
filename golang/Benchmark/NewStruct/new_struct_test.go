package NewStruct

import (
	"testing"
)

/*
	测试golang new 一个struct对性能及内存的损耗程度
*/

func BenchmarkNewStruct1(b *testing.B) {
	for i := 0; i < 1000000000; i++ {
		GetConfig(true)
	}
}

func BenchmarkNewStruct2(b *testing.B) {
	for i := 0; i < 1000000000; i++ {
		GetConfig(false)
	}
}
