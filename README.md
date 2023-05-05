# Golang 结构体转换

此工具用于将两个相似的结构体，a 值赋予 b 值。

结构体属性仅支持 Go 语言的基本类型
不支持递归比较嵌套结构体，匿名结构体

## 快速开始
安装
```bash
go install github.com/ixugo/atob
```
使用
```bash
# 查看版本号
atob -v
# 使用
atob
```

说明:

1. 输入完结构体，请按两次回车，表示输入完成
2. 当两个结构体使用了不同的属性名时，请在 a 结构体中使用 atob 标签

![demo](http://img.golang.space/img-1683297297032.png)