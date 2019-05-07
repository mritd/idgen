package test

import (
	"fmt"
	"testing"

	"github.com/mritd/idgen/utils"
)

func Test_GenFixedLengthChineseChars(t *testing.T) {
	fmt.Println(utils.GenFixedLengthChineseChars(10))
}

func Test_GenRandomLengthChineseChars(t *testing.T) {
	fmt.Println(utils.GenRandomLengthChineseChars(0, 10))
}

// 开发时 CardBin 生成测试
//func Test_CreateCardBinMetadata(t *testing.T) {
//	utils.CreateCardBinMetadata()
//}

func Benchmark_GenFixedLengthChineseChars(t *testing.B) {
	for i := 0; i < t.N; i++ {
		utils.GenFixedLengthChineseChars(10)
	}
}

func Benchmark_GenRandomLengthChineseChars(t *testing.B) {
	for i := 0; i < t.N; i++ {
		utils.GenRandomLengthChineseChars(0, 100)
	}
}
