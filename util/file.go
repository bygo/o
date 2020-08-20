package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func Replace(content string, replaces map[string]string) string {
	for old, new_ := range replaces {
		content = strings.Replace(content, old, new_, -1)
	}
	return content
}

func CheckExist(dir string) {
	_, err := os.Stat(dir)
	if err == nil || os.IsExist(err) {
		log.Fatal(fmt.Sprintf("Folder '%s' already exists", dir))
	}
}

func SnakeToCamel(name string) string {
	words := strings.Split(name, "_")
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	s := strings.Join(words, "")
	return s
}
