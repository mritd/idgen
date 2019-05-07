package generator

import (
	"strconv"

	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/utils"
)

// 随机省/城市
func GenProvinceAndCity() string {
	return metadata.ProvinceCity[utils.RandInt(0, len(metadata.ProvinceCity))]
}

// 随机地址生成
func AddrGenerate() string {
	return GenProvinceAndCity() +
		utils.GenRandomLengthChineseChars(2, 3) + "路" +
		strconv.Itoa(utils.RandInt(1, 8000)) + "号" +
		utils.GenRandomLengthChineseChars(2, 3) + "小区" +
		strconv.Itoa(utils.RandInt(1, 20)) + "单元" +
		strconv.Itoa(utils.RandInt(101, 2500)) + "室"
}
