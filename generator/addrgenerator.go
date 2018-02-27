package generator

import (
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/util"
	"strconv"
)

// 随机省/城市
func GenProvinceAndCity() string {
	return metadata.ProvinceCity[util.RandInt(0, len(metadata.ProvinceCity))]
}

// 随机地址生成
func AddrGenerate() string {
	return GenProvinceAndCity() +
		util.GenRandomLengthChineseChars(2, 3) + "路" +
		strconv.Itoa(util.RandInt(1, 8000)) + "号" +
		util.GenRandomLengthChineseChars(2, 3) + "小区" +
		strconv.Itoa(util.RandInt(1, 20)) + "单元" +
		strconv.Itoa(util.RandInt(101, 2500)) + "室"
}
