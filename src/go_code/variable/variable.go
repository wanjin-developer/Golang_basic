package main

import "fmt"

//演示变量

//定义全局变量
//第一种方式
var N1 = 1000
var N2 = 2000
var NAME = "cat"

//第二种方式，一次性声明
var (
	N3    = 3000
	N4    = 4000
	NAME2 = "jon"
)

func main() {
	//golang变量的使用方式
	//第一种：指定变量类型，声明后若不赋值，使用默认值
	//int 的默认值是
	var i int
	fmt.Println("i=", i)

	//第二种：根据值自行判定变量类型（类型推导）
	var num = 10.11
	fmt.Println("num=", num)

	//第三种：省略var，注意：=左侧的变量不应该是已经声明过的。否则会导致编译错误
	//等价于
	//var name string
	//name = "tom"
	name := "tom"
	fmt.Println("name=", name)

	//多变量声明
	//第一种方式
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//第二种方式
	var n4, name1, n5 = 100, "tom", 888
	fmt.Println("n4=", n4, "name1=", name1, "n5=", n5)

	//第三种方式：使用类型推导
	n6, name2, n7 := 100, "jack", 999
	fmt.Println("n6=", n6, "name2=", name2, "n7=", n7)

	//输出全局变量
	fmt.Println("N1=", N1, "N2=", N2, "NAME", NAME)
	fmt.Println("N3=", N3, "N4=", N4, "NAME2", NAME2)

	//+号使用
	var ch1 int = 100
	var ch2 int = 200
	fmt.Println(ch1+ch2)
	
	var s1 string = "hello"
	var s2 string = "world"
	var s3 string = s1 + s2
	fmt.Println("s3=",s3)

}

//变量使用注意事项：
//1:该区域的数据值可以在同一类型范围内不断变化
//var i int = 10
//i = 30
//i = 50
//fmt.Println("i=",i)
//i=1.2 此处报错，原因是不能改变数据类型

//2：变量在同一个作用域（一个函数或一个代码块）内不可以重复
//var i int = 59
//i := 100 此处报错，原因是不能重复

//3：变量=变量名+值+数据类型

//4：Golang如果没有初始值，编译器会使用默认值

//变量的声明：
//var 变量名 数据类型
//var a int 声明一个整型变量a
//var num1 float32 声明一个单精度浮点数num1

//初始化变量：
//声明时就给定值，可以省略数据类型
//var a int = 45  var b = 400

//给变量赋值
//var num int 
//num = 780

//程序中+号的使用
//左右两边是数值型时，则做加法运算

//左右是字符串时，则做字符串拼接