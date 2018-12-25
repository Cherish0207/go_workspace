package main
import "fmt"

//type SEX int
type Person struct {
	name string
	sex int
	addr string
}
type Student struct {
	p Person
	score int
}
func (s Student) PrintInfo() {
	fmt.Println(s)
}
func main() {
	stu := Student{{"crx",1,"sh"},39}
	stu.PrintInfo()
}