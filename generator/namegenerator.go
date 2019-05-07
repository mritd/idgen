package generator

import (
	"database/sql"

	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/utils"
)

// 生成姓名
func NameGenerate() string {
	if utils.DBExist() {
		var firstName string
		db, err := sql.Open("sqlite3", utils.DBPath())
		defer db.Close()
		utils.CheckAndExit(err)
		utils.CheckAndExit(db.QueryRow(utils.FirstNameSQL, utils.RandInt(0, utils.FirstNameSum)).Scan(&firstName))
		return metadata.LastName[utils.RandInt(0, len(metadata.LastName))] + firstName
	} else {
		return metadata.LastName[utils.RandInt(0, len(metadata.LastName))] + utils.GenRandomLengthChineseChars(1, 3)
	}
}
