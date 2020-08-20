package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func MKdir(paths ...string) {
	err := os.Mkdir(filepath.Join(paths...), 0755)
	if err != nil {
		panic(err)
	}
}

func WriteFile(content string, paths ...string) {
	var dir []string
	for _, path := range paths {
		dir = append(dir, path)
		MKdir(dir...)
	}
	ioutil.WriteFile(filepath.Join(paths...), []byte(content), 0755)
}

func Replace(content string, replaces map[string]string) string {
	for old, new_ := range replaces {
		content = strings.Replace(content, old, new_, -1)
	}
	return content
}
