package make

var schemaStub = `package DummyName

import "time"

type DummyName struct {
	ID        int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}`