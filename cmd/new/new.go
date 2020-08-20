package new

import (
	"github.com/temporaries/o/cmd"
	"github.com/temporaries/o/util/file"
	"log"
	"runtime"
	"strings"
)

var name cmd.Value

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
	name.Set(args[0])
	projectName := name.String()
	replaces["DummyProject"] = projectName

	//project exists
	file.CheckExist(projectName)

	//replace
	mainStub = file.Replace(mainStub, replaces)
	bootStub = file.Replace(bootStub, replaces)
	confAllStub = file.Replace(confAllStub, replaces)
	confDBStub = file.Replace(confDBStub, replaces)
	cn = file.Replace(cn, replaces)
	modStub = file.Replace(modStub, replaces)

	//write
	file.MKdir(projectName, "schema")
	file.WriteFile(mainStub, projectName, "main.go")
	file.WriteFile(bootStub, projectName, "boot", "boot.go")
	file.WriteFile(confAllStub, projectName, "config", "conf.go")
	file.WriteFile(confDBStub, projectName, "config", "db.go")
	file.WriteFile(cn, projectName, "config", "language", "cn.go")
	file.WriteFile(modStub, projectName, "go.mod")

	log.Print("Successfully Created!")
}
