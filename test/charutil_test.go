package test

import (
	"fmt"
	_ "github.com/mritd/idgen/statik"
	"github.com/mritd/idgen/util"
	"testing"
)

func Test_GenFixedLengthChineseChars(t *testing.T) {
	fmt.Println(util.GenFixedLengthChineseChars(10))
}

func Test_GenRandomLengthChineseChars(t *testing.T) {
	fmt.Println(util.GenRandomLengthChineseChars(0, 10))
}

// 开发时 CardBin 生成测试
//func Test_CreateCardBinMetadata(t *testing.T) {
//	util.CreateCardBinMetadata()
//}

func Benchmark_GenFixedLengthChineseChars(t *testing.B) {
	for i := 0; i < t.N; i++ {
		util.GenFixedLengthChineseChars(10)
	}
}

func Benchmark_GenRandomLengthChineseChars(t *testing.B) {
	for i := 0; i < t.N; i++ {
		util.GenRandomLengthChineseChars(0, 100)
	}
}
