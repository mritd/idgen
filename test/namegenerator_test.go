package test

import (
	"fmt"
	"github.com/mritd/idgen/generator"
	"testing"
)

func Test_NameGenerate(t *testing.T) {
	fmt.Println(generator.NameGenerate())
}

func Benchmark_NameGenerate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		generator.NameGenerate()
	}
}

//func Test_InitRealName(t *testing.T) {
//
//	os.Remove("../resources/data.db")
//	db, err := sql.Open("sqlite3", "../resources/data.db")
//	db.Exec("CREATE TABLE firstName( id INTEGER PRIMARY KEY AUTOINCREMENT , firstName VARCHAR(32) NULL );")
//
//	tx, err := db.Begin()
//	defer tx.Rollback()
//	insert, err := tx.Prepare("INSERT INTO firstName(firstName) values(?);")
//	defer insert.Close()
//	util.CheckAndExit(err)
//
//	f, err := os.Open("/Users/mritd/gopath/src/github.com/mritd/idgen/resources/fname")
//	b, err := ioutil.ReadAll(f)
//	s := strings.Fields(string(b))
//	for _, name := range s {
//		insert.Exec(name)
//	}
//	util.CheckAndExit(tx.Commit())
//
//}
