package core

import (
	"testing"
	"time"
)

type UserA struct {
	ID        int
	Name      string
	Weight    int
	CreatedAt time.Time `atob:"CreateTime"`
}
type UserB struct {
	ID         int
	Name       string
	Age        int
	CreateTime time.Time
}

func TestConvert(t *testing.T) {
	var a UserA
	var b UserB
	var temp Temp
	if err := Convert(&a, &b, &temp); err != nil {
		t.Fatal(err)
	}
	if temp.StructA != "UserA" {
		t.Fatal(`StructA != "UserA"`)
	}
	if temp.StructB != "UserB" {
		t.Fatal(`StructB != "UserB"`)
	}
	if temp.Name != "UserAToUserB" {
		t.Fatal(`Name != "UserAToUserB"`)
	}
	if len(temp.Fields) != 3 {
		t.Fatal(`len(temp.Fields) != 3`)
	}
}
