package main

import(
	"fmt"
)

func main(){
	//Golang没有专门字符类型,用byte保存
	//Golang字符串由字节组成
	var c1 byte = 'a'
	var c2 byte = '0'
	//直接输出byte值,会输出字符对应的ASCII码值
	fmt.Println("c1=",c1)
	fmt.Println("c2=",c2)
	//如果要输出对应字符,需要使用格式化输出
	fmt.Printf("c1=%c\n",c1)
	fmt.Printf("c2=%c\n",c2)


	//var c3 byte = '北' 这样会overflow溢出,因为byte类型存不了
	var c3 int = '北'
	fmt.Printf("c3=%c\n",c3)
	fmt.Printf("c3对应码值:%d\n",c3)
	//如果保存的字符在ASCII码中可以使用byte
	//如果码值大于255,可考虑使用int保存
	//如果需要字符格式化输出,要是用fmt.Printf("c3=%c\n",c3)
}