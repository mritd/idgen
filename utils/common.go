package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr/v2"

	"github.com/mitchellh/go-homedir"
)

func DBPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, ".idgen.db")
}

func ExportDB() {
	box := packr.New("resources", "../resources")
	bs, err := box.Find("idgen.db")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(DBPath(), bs, 0600)
	if err != nil {
		panic(err)
	}
}

func DBExist() bool {
	_, err := os.Stat(DBPath())
	return err == nil
}
