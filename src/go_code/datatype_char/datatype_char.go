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

	//Go允许使用转义字符'\'来将其后的字符转变为特殊字符型常量.例如 var c3 byte = '\n' , '\n'表示换行
	var c4 byte ='\n'
	fmt.Printf("c4=%c     mmm\n",c4)

	//可以直接给某个变量附一个数字,然后格式化输出时%c会输出改数字对应的Unicode字符
	var c5 int =22269
	fmt.Printf("c5对应的字符:%c\n",c5)

	//字符类型是可以运算的,相当于一个整数,因为他们都有对应的Unicode码
	var c6 = 10 + 'a'
	fmt.Println("c6=",c6)
	fmt.Printf("c6=%c\n",c6)

}

//字符类型使用细节:
//1:字符常量用单引号
//2:Go允许使用转义字符'\'来将其后的字符转变为特殊字符型常量.例如 var c3 byte = '\n' , '\n'表示换行
//3:Go字符使用UTF-8,英文1个字节,汉子3个字节
//4:GO中字符的本质是整数,直接输出时是该字符对应的UTF-8码值
//5:可以直接给某个变量附一个数字,然后格式化输出时%c会输出改数字对应的Unicode字符
//6:字符类型是可以运算的,相当于一个整数,因为他们都有对应的Unicode码

//字符类型本质:
//1:字符型存到计算机中,需要将字符对应的码值(整数)找出来
//存储:字符->对应码值->二进制->存储
//读取:二进制->码值->字符->读取
//2:字符和码值的对应关系是通过字符编码表决定的(规定好的)
//3:Go编码统一UTF-8