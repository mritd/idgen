package generator

import (
	"strconv"

	"github.com/etcd-io/bbolt"
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/utils"
)

// 生成姓名
func GetName() string {
	if utils.DBExist() {
		var firstName string
		db, err := bbolt.Open("idgen.db", 0600, nil)
		if err != nil {
			panic(err)
		}
		_ = db.View(func(tx *bbolt.Tx) error {
			firstName = string(tx.Bucket([]byte("firstName")).Get([]byte(strconv.Itoa(utils.RandInt(1, 786029)))))
			return nil
		})
		return metadata.LastName[utils.RandInt(0, len(metadata.LastName))] + firstName
	} else {
		return metadata.LastName[utils.RandInt(0, len(metadata.LastName))] + utils.GenRandomLengthChineseChars(1, 3)
	}
}
