package main

/*
	1. package main：声明hello.go所在的包，GO使用包来组织代码，一般一个文件夹即一个包，包内可以暴露类型或方法供其它包使用
	2. import "fmt"：fmt,Go语言标准库/包，用于标准输入输出
	3. func main：整个程序的入口
	4. fmt.Println("Hello World!"): Println,fmt包里的方法，用于打印输出
	5. go build 01_hello.go：编译成二进制可以执行程序
	6. go run 01_hello.go：编译并运行Go程序
*/

import "fmt"

func main() {
	fmt.Println("Hello World!") // Hello World!
}
