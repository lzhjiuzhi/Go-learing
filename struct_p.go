package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person //匿名字段继承 只有类型没有名字
	id     int
	addr   string
}

func main() {
	var s1 Student = Student{Person{"LIU", 'm', 10}, 1, "bj"} //顺序初始化
	fmt.Println(s1)
}
