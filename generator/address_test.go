package generator

import "testing"

func TestGetAddress(t *testing.T) {
	t.Log(GetAddress())
}

func BenchmarkGetAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetAddress())
	}
}
