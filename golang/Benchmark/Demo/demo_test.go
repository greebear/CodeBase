package Demo

import "testing"

/*
	参考连接：
	1. [Golang]Benchmark性能测试 https://blog.csdn.net/u013161278/article/details/117628603
*/

// Benchmark_test1
// 参与Benchmark基准性能测试的方法必须以Benchmark为前缀，例如BenchmarkABC()
func Benchmark_test1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		for i := 0; i < 10000; i++ {
			s = append(s, i)
		}
	}
}

func Benchmark_test2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		for i := 0; i < 10000; i++ {
			s = append(s, i)
		}
	}
}
