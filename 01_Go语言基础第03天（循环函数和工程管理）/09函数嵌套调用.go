package main

import "fmt"

func test2(a int, b int){
	sum:=a+b
	fmt.Println(sum)
}
func test1(a int , b int){
	test2(a,b)
}

//所有的函数都是全局函数 可以被项目中所有文件使用  在项目中函数名是唯一的
func main0901() {
	test1(10,20)

}


func test3(a ...int){
	//如果函数参数为不定参  传递方式为a[0:]...
	test4(a[0:]...)
	//test4(a[2:]...)
	//test4(a[:2]...)
	test4(a[1:3]...)
}
func main0902(){
	test3(1,2,3,4)
	test3(5,6,7,8)
}

//函数的作用域是全局的  可以在项目中任意文件使用 不需要在意先后顺序
func test4(b ...int){
	for i:=0;i<len(b);i++  {
		fmt.Println(i,b[i])
	}
}

func testq(args...int) {
	test4(args[0:]...)
}
func a(a ...int){
	//1 使用for循环
	for i:=0;i<len(a);i++ {
		fmt.Println(i,a[i])
	}
	//2 使用for range
	for i,v := range a{
		fmt.Println(i,v)
	}
	ab(a[0:]...)
	// 可以理解为 a[0:] 取到所有参数数组 ...展开数组
	// 可以使用 a[0:1] 这种写法
}
func ab(a ...int){
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[0:])
	fmt.Println(a[1:2])
	fmt.Println(a[:3])
}
func main() {
	a(1,3,4)
}