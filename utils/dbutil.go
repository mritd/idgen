package utils

import (
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/go-homedir"
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

func InitConfigDir() {
	home, err := homedir.Dir()
	CheckAndExit(err)
	ConfigDir := home + string(filepath.Separator) + ".idgen"
	_, err = os.Stat(ConfigDir)
	if err != nil {
		err = os.MkdirAll(ConfigDir, 0711)
		CheckAndExit(err)
	}
}

func DBExist() bool {
	_, err := os.Stat(DBPath())
	return err == nil
}
