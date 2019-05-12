package generator

import (
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/utils"
)

// 生成姓名
func GetName() string {
	return metadata.LastName[utils.RandInt(0, len(metadata.LastName))] + metadata.FirstName[utils.RandInt(0, len(metadata.LastName))]
}
