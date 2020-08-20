package cmd

import (
	"flag"
	"fmt"
	_ "github.com/temporaries/o/cmd/make"
	_ "github.com/temporaries/o/cmd/new"
	"os"
)

var Commands = map[string]*Command{}

type Command struct {
	Flag flag.FlagSet
	Run  func(c *Command, args []string)
}

func Usage() {
	os.Exit(2)
}

type Value string

func (v *Value) String() string {
	return fmt.Sprint(*v)
}

func (v *Value) Set(value string) error {
	*v = Value(value)
	return nil
}
