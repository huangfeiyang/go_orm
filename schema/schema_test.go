package schema

import (
	"testing"
	"geeorm/dialect"
)

type User struct {
	Name string `geeorm:"PRIMAY KEY"`
	Age int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("Fail to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMAY KEY" {
		t.Fatal("Fail to parse primary key")
	}
}