package test

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/mritd/idgen/utils"
)

func Test_DBPath(t *testing.T) {
	fmt.Println(utils.DBPath())
}

func Test_DBExist(t *testing.T) {
	fmt.Println(utils.DBExist())
}

func Test_InitRealName(t *testing.T) {

	t.Skip("跳过 DB 生成(仅用于本地测试)")

	os.Remove("../resources/data.db")
	db, err := sql.Open("sqlite3", "../resources/data.db")
	defer db.Close()
	db.Exec("CREATE TABLE firstName( id INTEGER PRIMARY KEY AUTOINCREMENT , firstName VARCHAR(32) NULL );")

	tx, err := db.Begin()
	defer tx.Rollback()
	insert, err := tx.Prepare("INSERT INTO firstName(firstName) values(?);")
	defer insert.Close()
	utils.CheckAndExit(err)

	f, err := os.Open("../resources/fname")
	b, err := ioutil.ReadAll(f)
	s := strings.Fields(string(b))
	for _, name := range s {
		insert.Exec(name)
	}
	utils.CheckAndExit(tx.Commit())

}

func Benchmark_DBPath(t *testing.B) {
	for i := 0; i < t.N; i++ {
		utils.DBPath()
	}
}

func Benchmark_DBExist(t *testing.B) {
	for i := 0; i < t.N; i++ {
		utils.DBExist()
	}
}
