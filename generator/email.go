package generator

import (
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/utils"
)

// 随机生成邮箱
func GetEmail() string {
	return utils.RandStr(8) + "@" + utils.RandStr(5) + metadata.DomainSuffix[utils.RandInt(0, len(metadata.DomainSuffix))]
}
