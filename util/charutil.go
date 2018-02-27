package util

import (
	"github.com/mritd/idgen/metadata"
	"math/rand"
	"time"
)

// 随机单个中文字符
func GenOneChineseChars() string {
	return string(rune(RandInt(19968, 40869)))
}

// 指定长度随机中文字符
func GenFixedLengthChineseChars(length int) string {
	strRune := make([]rune, length)
	for i := range strRune {
		strRune[i] = rune(RandInt(19968, 40869))
	}
	return string(strRune)
}

// 指定范围随机中文字符
func GenRandomLengthChineseChars(start, end int) string {
	return GenFixedLengthChineseChars(RandInt(start, end))
}

// 随机单个复杂中文字符
func GetOneOddChar() string{
	rand.Seed(time.Now().UnixNano())
	return string([]rune(metadata.ODD_CHINESE_CHARS)[rand.Intn(len([]rune(metadata.ODD_CHINESE_CHARS)))])
}

// 指定范围随机 int
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// 指定范围随机 int64
func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}
