package core

import (
	"fmt"
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

func TestCompareStructFields(t *testing.T) {
	const a = "package abc\ntype UserA struct {\n" +
		"\tID        int\n" +
		"\tName      string\n" +
		"\tWeight    int\n" +
		"\tCreatedAt time.Time `atob:\"CreateTime\"`\n" +
		"}"

	const b = `
package abc
type UserB struct {
	ID         int
	Name       string
	Age        int
	CreateTime time.Time
}`
	var temp Temp
	if err := CompareStructFields(a, b, &temp); err != nil {
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

// TestComplexCompareStructFields 不支持嵌套
func TestComplexCompareStructFields(t *testing.T) {
	const a = "package abc\ntype UserA struct {\n" +
		"\tID        int\n" +
		"\tName      string\n" +
		"\tWeight    int\n" +
		"\tCreatedAt time.Time `atob:\"CreateTime\"`\n" +
		`Device struct {
			Name string
		}
		Role int
		` +
		"}"

	const b = `
package abc
type UserB struct {
	ID         int
	Name       string
	Age        int
	CreateTime time.Time
	Device struct {
		Name string
	}
	Role string
}`
	var temp Temp
	if err := CompareStructFields(a, b, &temp); err != nil {
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

	for _, v := range temp.Fields {
		fmt.Println(v)
	}
}
