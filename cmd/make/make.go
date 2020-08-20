package make

import (
	"github.com/temporaries/o/cmd"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var name cmd.Value

var CmdMake = &cmd.Command{
	Run: Run,
}

var replaces = map[string]string{
	"DummyName": "",
}

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

	mkdir("schema", schemaName)
	writeFile(schemaStub, "schema", schemaName, schemaName+".go")
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
