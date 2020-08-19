package new

import (
	"fmt"
	"github.com/temporaries/o/cmd"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var name cmd.Value

var CmdNew = &cmd.Command{
	Run: Run,
}

var replaces = map[string]string{
	"DummyProject":     name.String(),
	"DummyGoVersion":   strings.TrimLeft(runtime.Version(), "go"),
	"DummyLureVersion": "0.0.2",
}

func init() {
	//CmdNew.Flag.Var(&name, "name", "project name")
	cmd.Commands["new"] = CmdNew
}

func Run(c *cmd.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("Please enter the project name")
	}

	name.Set(args[0])
	project := name.String()
	_, err := os.Stat(project)
	if err == nil || os.IsExist(err) {
		log.Fatal(fmt.Sprintf("Folder '%s' already exists", project))
	}

	log.Print("Creating...")

	mkdir()
	mkdir("boot")
	mkdir("schema")
	mkdir("config")
	mkdir("config", "language")

	writeFile(mainStub, "main.go")
	writeFile(bootStub, "boot", "boot.go")
	writeFile(confAllStub, "config", "conf.go")
	writeFile(confDBStub, "config", "db.go")
	writeFile(cn, "config", "language", "cn.go")
	writeFile(modStub, "go.mod")

	log.Print("Successfully Created!")
}

func mkdir(paths ...string) {
	err := os.Mkdir(filepath.Join(append([]string{name.String()}, paths...)...), 0755)
	if err != nil {
		panic(err)
	}
}

func writeFile(content string, path ...string) {
	for old, new_ := range replaces {
		content = strings.Replace(content, old, new_, -1)
	}

	paths := append([]string{name.String()}, path...)

	filename := filepath.Join(paths...)

	ioutil.WriteFile(filename, []byte(content), 0755)
}
