package main

import "fmt"

// 计算斐波那契数
func fib(n int) ([]int, int) {
	//if n <= 0 {
	//	return [], 0
	//}
	var result = make([]int, n+1)
	result[0] = 0
	result[1] = 1
	if n <= 1 {
		return result, result[n]
	}
	for i:=2;i<=n;i++ {
		result[i] = result[i-2] + result[i-1]
	}
	//fmt.Printf("%T", result)
	return result, result[n]
}
func main_fib() {
	 result, res := fib(10)
	fmt.Println(result)
	fmt.Println(res)
}
func main() {

}
