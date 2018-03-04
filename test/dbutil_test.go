package test

import (
	"database/sql"
	"fmt"
	"github.com/mritd/idgen/util"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_DBPath(t *testing.T) {
	fmt.Println(util.DBPath())
}

func Test_DBExist(t *testing.T) {
	fmt.Println(util.DBExist())
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
	util.CheckAndExit(err)

	f, err := os.Open("../resources/fname")
	b, err := ioutil.ReadAll(f)
	s := strings.Fields(string(b))
	for _, name := range s {
		insert.Exec(name)
	}
	util.CheckAndExit(tx.Commit())

}

func Benchmark_DBPath(t *testing.B) {
	for i := 0; i < t.N; i++ {
		util.DBPath()
	}
}

func Benchmark_DBExist(t *testing.B) {
	for i := 0; i < t.N; i++ {
		util.DBExist()
	}
}
