package main
import "testing"

func BenchmarkLoopCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loopCalc(1000)
	}
}

func BenchmarkParallelCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallelCalc(1000)
	}
}

