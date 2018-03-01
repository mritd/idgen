package test

import (
	"fmt"
	"github.com/mritd/idgen/util"
	"testing"
	_ "github.com/mritd/idgen/statik"
)

func Test_GenOneChineseChars(t *testing.T) {
	fmt.Println(util.GenOneChineseChars())
}

func Test_GenFixedLengthChineseChars(t *testing.T) {
	fmt.Println(util.GenFixedLengthChineseChars(100))
}

func Test_GenRandomLengthChineseChars(t *testing.T) {
	fmt.Println(util.GenRandomLengthChineseChars(0, 100))
}

func Test_GenFirstName(t *testing.T) {
	util.GenFirstName()
}

// 开发时 CardBin 生成测试
//func Test_CreateCardBinMetadata(t *testing.T) {
//	util.CreateCardBinMetadata()
//}

func Benchmark_GenOneChineseChars(t *testing.B) {
	for i := 0; i < t.N; i++ {
		util.GenOneChineseChars()
	}
}

func Benchmark_GenFixedLengthChineseChars(t *testing.B) {
	for i := 0; i < t.N; i++ {
		util.GenFixedLengthChineseChars(100)
	}
}

func Benchmark_GenRandomLengthChineseChars(t *testing.B) {
	for i := 0; i < t.N; i++ {
		util.GenRandomLengthChineseChars(0, 100)
	}
}
