package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func MKdir(paths ...string) {
	err := os.MkdirAll(filepath.Join(paths...), 0755)
	if err != nil {
		panic(err)
	}
}

func WriteFile(content string, paths ...string) {
	l := len(paths)
	if 1 < l {
		MKdir(paths[:l-1]...)
	}
	ioutil.WriteFile(filepath.Join(paths...), []byte(content), 0755)
}

func IsExist(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil || os.IsExist(err)
}
