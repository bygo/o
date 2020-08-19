package new

var bootStub = `package boot

import (
	"github.com/temporaries/orc/app"
	"github.com/temporaries/orc/conf"
	"github.com/temporaries/orc/schema"
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
