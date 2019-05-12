package generator

import "testing"

func TestGetMobile(t *testing.T) {
	t.Log(GetMobile())
}

func BenchmarkGetMobile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetMobile())
	}
}
