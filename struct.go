package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func test01(s Student){
	s.id = 666
	fmt.Println("s.id = " ,s.id)
	//值传递

func test02(p *Student){
	p.id = 666
}
}
func main() {
	//顺序初始化 每个成员都要初始化
	var s1 Student = Student{1, "mike", 'm', 18, "asdasd"}
	fmt.Println(s1)
	//指定初始化 其他为0
	s2 := Student{name: "mike", addr: "QINGDAO"}
	fmt.Println(s2)
	//操作成员用.
	var s3 Student
	s3.id = 1
	s3.name = "ASD"
	s3.sex = 'm'
	s3.age = 11
	fmt.Println(s3)
	test01(s3)
	test02(&s3)

}
