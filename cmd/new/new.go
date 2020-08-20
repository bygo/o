package new

import (
	"github.com/temporaries/o/cmd"
	"github.com/temporaries/o/util"
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
	util.CheckExist(projectName)

	log.Print("Creating...")

	//replace
	mainStub = util.Replace(mainStub, replaces)
	bootStub = util.Replace(bootStub, replaces)
	confAllStub = util.Replace(confAllStub, replaces)
	confDBStub = util.Replace(confDBStub, replaces)
	cn = util.Replace(cn, replaces)
	modStub = util.Replace(modStub, replaces)

	//write
	util.MKdir(projectName, "schema")
	util.WriteFile(mainStub, projectName, "main.go")
	util.WriteFile(bootStub, projectName, "boot", "boot.go")
	util.WriteFile(confAllStub, projectName, "config", "conf.go")
	util.WriteFile(confDBStub, projectName, "config", "db.go")
	util.WriteFile(cn, projectName, "config", "language", "cn.go")
	util.WriteFile(modStub, "go.mod")

	log.Print("Successfully Created!")
}
