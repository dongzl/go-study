package bench_test

import (
	bench "github.com/arana-db/arana/bench_test"
	"testing"
	"time"
)

// https://geektutu.com/post/hpg-benchmark.html
func BenchmarkFib(b *testing.B) {
	time.Sleep(time.Second * 3) // 模拟耗时准备任务
	b.ResetTimer()              // 重置定时器
	for n := 0; n < b.N; n++ {
		bench.Fib(30)
	}
}
