package generator

import "testing"

func TestGetName(t *testing.T) {
	t.Log(GetName())
}

func BenchmarkGetName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetName())
	}
}
