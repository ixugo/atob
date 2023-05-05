package core

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

func CompareStructFields(a, b string, bean *Temp) error {
	fset := token.NewFileSet()
	aFile, err := parser.ParseFile(fset, "", a, parser.ParseComments)
	if err != nil {
		return err
	}
	bFile, err := parser.ParseFile(fset, "", b, parser.ParseComments)
	if err != nil {
		return err
	}

	aStruct, aStrcutName := findFirstStruct(aFile)
	bStruct, bStructName := findFirstStruct(bFile)

	cache := structFields(aStruct)

	result := make([]string, 0, 10)
	for _, bField := range bStruct.Fields.List {
		bFieldName := bField.Names[0].Name
		if aField, ok := cache[bFieldName]; ok {
			aFieldName := aField.Names[0].Name
			var text string
			if fmt.Sprint(aField.Type) == fmt.Sprint(bField.Type) {
				text = fmt.Sprintf("\t%s.%s = %s.%s \n", "b", bFieldName, "a", aFieldName)
			} else {
				text = fmt.Sprintf("\t%s.%s = %s.%s // 类型疑似不同，请手动调整\n", "b", bFieldName, "a", aFieldName)
			}
			result = append(result, text)
			// fmt.Printf("%s.%s = %s.%s\n", aStrcutName, bFieldName, bStructName, aFieldName)
		}
	}

	if len(result) > 0 {
		result[len(result)-1] = strings.TrimRight(result[len(result)-1], "\n")
	}

	bean.Fields = result
	bean.StructA = aStrcutName
	bean.StructB = bStructName
	bean.Name = fmt.Sprintf("%sTo%s", bean.StructA, bean.StructB)
	return nil
}

func findFirstStruct(file *ast.File) (*ast.StructType, string) {
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			if typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec); ok {
				if structType, ok := typeSpec.Type.(*ast.StructType); ok {
					return structType, typeSpec.Name.Name
				}
			}
		}
	}
	return nil, ""
}

func structFields(structType *ast.StructType) map[string]*ast.Field {
	fields := make(map[string]*ast.Field)
	for _, field := range structType.Fields.List {
		var fieldName string
		if len(field.Names) > 0 {
			fieldName = field.Names[0].Name
		}
		if field.Tag != nil {
			tag := field.Tag.Value
			fieldName = reflect.StructTag(tag[1 : len(tag)-1]).Get("atob")
		}
		if fieldName != "" {
			fields[fieldName] = field
		}
	}
	return fields
}
