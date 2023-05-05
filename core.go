package core

import (
	"fmt"
	"reflect"
	"strings"
)

type Temp struct {
	Name    string
	StructA string
	StructB string
	Fields  []string
}

func Convert(a, b any, bean *Temp) error {
	// a,b 是结构体
	// 通过反射将 a 和 b 相同的字段进行转换
	aType := reflect.TypeOf(a).Elem()
	bType := reflect.TypeOf(b).Elem()

	result := make([]string, 0, 10)
	for i := 0; i < bType.NumField(); i++ {
		fieldB := bType.Field(i)
		fieldA, ok := aType.FieldByName(fieldB.Name)
		if !ok {
			idx, exist := IndexOf(aType, fieldB.Name)
			if !exist {
				continue
			}
			fieldA = aType.Field(idx)
		}
		if fieldB.Type != fieldA.Type {
			continue
		}
		result = append(result, fmt.Sprintf("\t%s.%s = %s.%s\n", "b", fieldB.Name, "a", fieldA.Name))
	}

	// 去掉最后一个元素的换行
	result[len(result)-1] = strings.TrimRight(result[len(result)-1], "\n")
	bean.Fields = result
	bean.StructA = aType.Name()
	bean.StructB = bType.Name()
	bean.Name = fmt.Sprintf("%sTo%s", bean.StructA, bean.StructB)
	return nil
}

func IndexOf(rt reflect.Type, name string) (int, bool) {
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag.Get("atob")
		if tag == name {
			return i, true
		}
	}
	return -1, false
}
