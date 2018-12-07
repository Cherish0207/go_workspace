package main

import "fmt"

func main0201() {
	var d1 byte
	fmt.Printf("%c\n",d1) // (空)
	fmt.Printf("%T\n",d1) // unit8
	fmt.Println(d1) // 0
	fmt.Println()

	var d2 byte ='A'
	fmt.Printf("%c\n",d2) // A
	fmt.Printf("%T\n",d2) // unit8
	fmt.Println(d2) // 65
	fmt.Println()

	d3:='A'
	fmt.Printf("%c\n",d3) // A
	fmt.Printf("%T\n",d3) // int32
	fmt.Println(d3) // 65
	fmt.Println()

	var s string ="hello"
	fmt.Printf("%s\n",s)
	fmt.Printf("%T\n",s)
	//%% 会打印一个%
	fmt.Printf("35%%")
}
func main(){
	//计算机能够识别的进制  二进制 八进制 十进制 十六进制

	a:=123//十进制数据
	b:=0123//八进制数据 以0开头的数据是八进制
	c:=0xabc//十六进制  以0x开头的数据是十六进制
	//go语言中不能直接表示二进制数据

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//%b 占位符 表示输出一个二进制数据
	fmt.Printf("二进制值为：%b\n",a)
	fmt.Printf("二进制值为：%b\n",b)
	fmt.Printf("二进制值为：%b\n",c)
	//%o 占位符 表示输出一个八进制数据
	fmt.Printf("八进制值为：%o\n",a)
	fmt.Printf("八进制值为：%o\n",b)
	fmt.Printf("八进制值为：%o\n",c)
	//%x %X 占位符 表示输出一个十六进制数据
	fmt.Printf("十六进制值为：%X\n",a)
	fmt.Printf("十六进制值为：%X\n",b)
	fmt.Printf("十六进制值为：%X\n",c)
	fmt.Printf("十六进制值为：%x\n",a)
	fmt.Printf("十六进制值为：%x\n",b)
	fmt.Printf("十六进制值为：%x\n", c)
	s:=' '
	fmt.Printf("%T",s)
}