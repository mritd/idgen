package generator

import "testing"

func TestGetBank(t *testing.T) {
	t.Log(GetBank())
}

func BenchmarkGetBank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetBank())
	}
}
