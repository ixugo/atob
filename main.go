package main

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/ixugo/atob/core"
)

const pkName = "package atob\n\n"

//go:embed func.temp
var tempStr embed.FS

// input 从终端获取输入，遇到空串时结束
func input(desc string, value *string) error {
	fmt.Println(desc)
	scanner := bufio.NewScanner(os.Stdin)
	var result string
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" && result != "" {
			break
		}
		result += text + "\n"
	}
	*value = result
	return scanner.Err()
}

func recordErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-v" {
		fmt.Println("0.1.2")
		return
	}
	var a, b string
	recordErr(input("请输入 a 结构体(输入结束按 2 下回车):", &a))
	recordErr(input("请输入 b 结构体(输入结束按 2 下回车):", &b))

	fmt.Println("正在生成函数...")
	var temp core.Temp
	recordErr(core.CompareStructFields(pkName+a, pkName+b, &temp))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>")

	t := template.Must(template.ParseFS(tempStr, "func.temp"))
	recordErr(t.Execute(os.Stdout, temp))
}
