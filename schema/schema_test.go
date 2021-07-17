package schema

import (
	"github.com/KarlvenK/korm/dialect"
	"testing"
)

type User struct {
	Name string `korm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	scheam := Parse(&User{}, TestDial)
	if scheam.Name != "User" || len(scheam.FieldNames) != 2 {
		t.Fatal("failed to parese User struct")
	}
	if scheam.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
