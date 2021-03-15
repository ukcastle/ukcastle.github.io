package pop_test

import (
	"ch2/pop/pop"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pop.PopCount(101)
	}
}
