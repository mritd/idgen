package generator

import "testing"

func TestLUHNProcess(t *testing.T) {
	t.Log(LUHNProcess("623190380371814"))
}

func TestGetBank(t *testing.T) {
	t.Log(GetBank())
}

func BenchmarkLUHNProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(LUHNProcess("623190380371814"))
	}
}

func BenchmarkGetBank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetBank())
	}
}
