package bench_test

import "testing"

// https://geektutu.com/post/hpg-benchmark.html
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}
