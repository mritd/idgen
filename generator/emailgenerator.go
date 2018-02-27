package generator

import (
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/util"
)

// 随机生成邮箱
func EmailGenerate() string {
	return util.RandStr(8) + "@" + util.RandStr(5) + metadata.DomainSuffix[util.RandInt(0, len(metadata.DomainSuffix))]
}
