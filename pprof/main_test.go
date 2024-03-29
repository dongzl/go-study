package main

import (
	"github.com/henrygas/go-pprof-example/data"
	"testing"
)

const url = "https://github.com/EDDYCJY"

func TestAdd(t *testing.T) {
	s := data.Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data.Add(url)
	}
}
