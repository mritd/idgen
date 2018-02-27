package test

import (
	"fmt"
	"github.com/mritd/idgen/generator"
	"testing"
)

func Test_EmailGenerate(t *testing.T) {
	fmt.Println(generator.EmailGenerate())
}

func Benchmark_EmailGenerate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.EmailGenerate()
	}
}
