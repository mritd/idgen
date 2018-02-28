package test

import (
	"fmt"
	"github.com/mritd/idgen/generator"
	"testing"
)

func Test_BankGenerate(t *testing.T) {
	fmt.Println(generator.BankGenerate())
}

func Benchmark_BankGenerate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.MobileGenerate()
	}
}
