package new

import (
	"fmt"
	"github.com/bygo/o/cmd"
	"github.com/bygo/orc/util/file"
	"github.com/bygo/orc/util/str"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//var name cmd.Value

var CmdNew = &cmd.Command{
	Run: Run,
}

var replaces = map[string]string{
	"DummyProject":    "",
	"DummyGoVersion":  strings.TrimLeft(runtime.Version(), "go"),
	"DummyORCVersion": "0.0.2",
}

func init() {
	//CmdNew.Flag.Var(&name, "name", "project name")
	cmd.Commands["new"] = CmdNew
}

func Run(c *cmd.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("Please enter the project name")
	}

	//params
	//name.Set(args[0])
	projectName := args[0]
	replaces["DummyProject"] = projectName

	//project exists
	if file.IsExist(projectName) {
		log.Fatal(fmt.Sprintf("Folder '%s' already exists", projectName))
	}

	//replace
	mainStub = str.ReplaceMap(mainStub, replaces)
	bootStub = str.ReplaceMap(bootStub, replaces)
	confAllStub = str.ReplaceMap(confAllStub, replaces)
	confDBStub = str.ReplaceMap(confDBStub, replaces)
	cn = str.ReplaceMap(cn, replaces)
	modStub = str.ReplaceMap(modStub, replaces)

	//write
	schemaDir := filepath.Join(projectName, "schema")
	err := os.MkdirAll(schemaDir, 0755)
	if err != nil {
		panic(err)
	}

	file.WriteFileIfNotExist(filepath.Join(projectName, "main.go"), []byte(mainStub))
	file.WriteFileIfNotExist(filepath.Join(projectName, "boot", "boot.go"), []byte(bootStub))
	file.WriteFileIfNotExist(filepath.Join(projectName, "config", "conf.go"), []byte(confAllStub))
	file.WriteFileIfNotExist(filepath.Join(projectName, "config", "db.go"), []byte(confDBStub))
	file.WriteFileIfNotExist(filepath.Join(projectName, "config", "language", "cn.go"), []byte(cn))
	file.WriteFileIfNotExist(filepath.Join(projectName, "go.mod"), []byte(modStub))

	log.Print("Successfully Created!")
}
