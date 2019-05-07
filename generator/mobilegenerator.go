package generator

import (
	"fmt"

	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/utils"
)

// 随机生成手机号
func MobileGenerate() string {
	return metadata.MobilePrefix[utils.RandInt(0, len(metadata.MobilePrefix))] + fmt.Sprintf("%0*d", 8, utils.RandInt(0, 100000000))
}
