package main

import(
	"fmt"
)

//基本数据类型转换
func main(){
	//Go在不同类型的变量之间赋值需要显示转换,也就是说Golang中数据类型不能自动转换
	//基本语法
	//T:就是数据类型,比如int32,int64,float32等
	//v:是需要转换的变量
	var i int = 100
	var f float64 = float64(i)
	var u uint8 = uint8(f)
	fmt.Printf("i=%v,f=%v,u=%v\n",i,f,u)

	//从低精度->高精度也要进行转换
	var n1 int8 = 125
	var n2 int64 = int64(n1)
	fmt.Printf("n1=%v,n2=%v\n",n1,n2)

	//变量转换注意细节:
	//Go中,数据类型可以从表示范围小->表示范围大,也可以从范围大->范围小

	//被转换的是变量存储的数据(值),变量本身的数据类型并没有变化
	//i是int32
	fmt.Printf("i 的数据类型是:%T\n",i)

	//在转换中,从int64转成int8,编译时不会报错,只是转换结果是按照溢出处理,和希望的结果不一样
	var num1 int64 = 9999999
	var num2 int8 = int8(num1)
	fmt.Printf("num1=%v,num2=%v\n",num1,num2)

	var n3 int32 = 12
	var n4 int64
	var n5 int8 
	//int32->int64错误,需要转换n1类型
	n4 = int64(n3) + 20 
	//int32->int8错误
	n5 = int8(n3) + 20
	fmt.Printf("n4=%v,n5=%v",n4,n5)


}