package test

import (
	"fmt"
	"github.com/mritd/idgen/generator"
	"testing"
)

func Test_GenerateIssueOrg(t *testing.T) {
	fmt.Println(generator.GenerateIssueOrg())
}

func Test_GenerateValidPeriod(t *testing.T) {
	fmt.Println(generator.GenerateValidPeriod())
}

func Test_IDCardGenerate(t *testing.T) {
	fmt.Println(generator.IDCardGenerate())
}

func Benchmark_GenerateIssueOrg(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.GenerateIssueOrg()
	}
}

func Benchmark_GenerateValidPeriod(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.GenerateValidPeriod()
	}
}

func Benchmark_IDCardGenerate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.IDCardGenerate()
	}
}
