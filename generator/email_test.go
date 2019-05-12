package generator

import "testing"

func TestGetEmail(t *testing.T) {
	t.Log(GetEmail())
}

func BenchmarkGetEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetEmail())
	}
}
