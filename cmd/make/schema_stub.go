package make

var schemaStub = `package DummyPackageName

import "time"

type DummyStructName struct {
	ID        int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}`