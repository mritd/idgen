package generator

import "testing"

func TestGetProvinceAndCity(t *testing.T) {
	t.Log(GetProvinceAndCity())
}

func TestGetAddress(t *testing.T) {
	t.Log(GetAddress())
}

func BenchmarkGetProvinceAndCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetProvinceAndCity())
	}
}

func BenchmarkGetAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetAddress())
	}
}
