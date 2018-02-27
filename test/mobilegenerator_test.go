package test

import (
	"fmt"
	"github.com/mritd/idgen/generator"
	"testing"
)

func Test_MobileGenerate(t *testing.T) {
	fmt.Println(generator.MobileGenerate())
}

func Benchmark_MobileGenerate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.MobileGenerate()
	}
}
