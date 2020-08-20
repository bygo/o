package make

import (
	"github.com/temporaries/o/cmd"
	"github.com/temporaries/orc/util/file"
	"github.com/temporaries/orc/util/str"

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

	for _, name := range args {
		schemaName := str.CamelToSnake(name)
		structName := str.SnakeToBigCamel(schemaName)
		replaces["DummyPackageName"] = schemaName
		replaces["DummyStructName"] = structName

		schema := str.ReplaceMap(schemaStub, replaces)
		file.WriteFile(schema, "schema", schemaName, "schema.go")
		log.Print(name + " Created!")
	}
}
