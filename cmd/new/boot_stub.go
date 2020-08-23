package new

var bootStub = `package boot

import (
	"github.com/bygo/orc/app"
	"github.com/bygo/orc/conf"
	"github.com/bygo/orc/schema"
	"DummyProject/config"
)

func Start() {
	app.DeepRegister(
		Config,
	)
	app.Boot()

	schema.Migrate("./schema")
}

func Config() *conf.All {
	return &config.Config
}
`
