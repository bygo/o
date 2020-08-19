package new

var confAllStub = `package config

import "github.com/temporaries/orc/conf"

var Config = conf.All{
	DB: db,
}
`

var confDBStub = `package config

import (
	"github.com/temporaries/orc/conf"
	"DummyProject/config/language"
)

var db = conf.DB{
	Engine:   "mysql",
	Language: language.CN,
	Mysql: conf.Mysql{
		Host:      "127.0.0.1",
		Port:      "3306",
		Database:  "orc",
		Username:  "root",
		Password:  "root",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_unicode_ci",
		Prefix:    "",
	},
}
`

var cn = `
package language

var CN = map[string]string{
	"ID":        "ID",
	"Name":      "名字",
	"Title":     "标题",
	"Bio":       "简介",
	"Mobile":    "手机",
	"Status":    "状态",
	"Amount":    "数额",
	"CreatedAt": "创建时间",
	"UpdatedAt": "更新时间",
	"DeletedAt": "删除时间",
}
`
