package test

import (
	"fmt"
	"github.com/mritd/idgen/generator"
	"testing"
)

func Test_GenProvinceAndCity(t *testing.T) {
	fmt.Println(generator.GenProvinceAndCity())
}

func Test_AddrGenerate(t *testing.T) {
	fmt.Println(generator.AddrGenerate())
}

func Benchmark_GenProvinceAndCity(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.GenProvinceAndCity()
	}
}

func Benchmark_AddrGenerate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.AddrGenerate()
	}
}
