package make

import (
	"github.com/temporaries/o/cmd"
	"github.com/temporaries/o/util"
	"log"
)

var name cmd.Value

var CmdMake = &cmd.Command{
	Run: Run,
}

var replaces = map[string]string{}

func init() {
	cmd.Commands["make"] = CmdMake
}

func Run(c *cmd.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("Please enter the schema name")
	}

	name.Set(args[0])
	schemaName := name.String()
	replaces["DummyName"] = schemaName

	log.Print("Creating...")

	schemaStub = util.Replace(schemaStub, replaces)
	util.WriteFile(schemaStub, "schema", schemaName, schemaName+".go")

	log.Print("Successfully Created!")
}
