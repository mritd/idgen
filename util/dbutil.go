package util

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/go-homedir"
	"github.com/mritd/idgen/metadata"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	FirstNameSQL = "SELECT firstName FROM firstName WHERE id = ?"
	FirstNameSum = 786029
)

func DBPath() string {
	home, err := homedir.Dir()
	CheckAndExit(err)
	return home + string(filepath.Separator) + ".idgen/data.db"
}

func InitConfigDir(){
    home, err := homedir.Dir()
    CheckAndExit(err)
    ConfigDir :=home + string(filepath.Separator) + ".idgen"
    _,err=os.Stat(ConfigDir)
    if err!=nil {
        err = os.MkdirAll(ConfigDir,0711)
        CheckAndExit(err)
    }
}

func DBExist() bool {
	_, err := os.Stat(DBPath())
	return err == nil
}

func DB() *sql.DB {
	db, err := sql.Open("sqlite3", DBPath())
	CheckAndExit(err)
	return db
}

func InitDB() {
    InitConfigDir()
	res, err := http.Get(metadata.DB_DOWNLOAD_URL)
	CheckAndExit(err)
	db, err := os.OpenFile(DBPath(), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	CheckAndExit(err)
	io.Copy(db, res.Body)
}
