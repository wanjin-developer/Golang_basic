package main

import (
	"fmt"
	"strconv"
)

//string转基本数据类型
func main() {
	//使用strconv包的函数

	//func ParseBool(str string) (value bool, err error)
	//返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；
	//否则返回错误。
	//1:strconv.ParseBool(str) 函数会返回两个值(value bool ,err error)
	//2:因为我只想获取到value bool,不想获取err,所以使用_忽略
	var str3 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str3)
	fmt.Printf("b type %T,b value %v\n", b, b)

	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//返回字符串表示的整数值，接受正负号。
	//base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	//返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
	var str1 string = "100"
	var n1 int64
	n1, _ = strconv.ParseInt(str1, 0, 32)
	fmt.Printf("n1 type %T,n1 value %v\n", n1, n1)

	//func ParseFloat(s string, bitSize int) (f float64, err error)
	//解析一个表示浮点数的字符串并返回其值。
	//如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。
	//bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
	//返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；
	//结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
	var str2 string = "3.145"
	var n2 float64
	n2,_ = strconv.ParseFloat(str2, 64)
	fmt.Printf("n2 type %T,n2 value %v\n", n2, n2)

	//注意事项:
	//在将string类型转换成基本数据类型时,要确保string类型能够转成有效的数据.
	//比如:可以吧"123"转成一个整数,但是不能把"hello"转成一个整数,如果这样做,Golang直接将其转成0,其他类型也一致
	var str4 string = "hello"
	var num4 int64
	num4,_ = strconv.ParseInt(str4,0,64)
	fmt.Printf("num4 type %T,num4 value %v",num4,num4)
}