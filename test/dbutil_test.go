package test

import (
	"fmt"
	"github.com/mritd/idgen/util"
	"testing"
)

func Test_DBPath(t *testing.T) {
	fmt.Println(util.DBPath())
}

func Test_DBExist(t *testing.T) {
	fmt.Println(util.DBExist())
}

func Benchmark_DBPath(t *testing.B) {
	for i := 0; i < t.N; i++ {
		util.DBPath()
	}
}

func Benchmark_DBExist(t *testing.B) {
	for i := 0; i < t.N; i++ {
		util.DBExist()
	}
}
