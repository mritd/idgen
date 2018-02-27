package generator

import (
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/util"
)

// 生成姓名
func NameGenerate() string {
	return metadata.LastName[util.RandInt(0, len(metadata.LastName))] + util.GenRandomLengthChineseChars(1, 3)
}

// 生成带有生僻字的姓名
func NameGenerateOdd() string {
	return metadata.LastName[util.RandInt(0, len(metadata.LastName))] + util.OneOddChar()
}
