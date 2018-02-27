package test

import (
	"fmt"
	"github.com/mritd/idgen/generator"
	"testing"
)

func Test_NameGenerate(t *testing.T) {
	fmt.Println(generator.NameGenerate())
}

func Test_NameGenerateOdd(t *testing.T) {
	fmt.Println(generator.NameGenerateOdd())
}

func Benchmark_NameGenerate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.NameGenerate()
	}
}

func Benchmark_NameGenerateOdd(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.NameGenerateOdd()
	}
}
