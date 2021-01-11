package main

import(
	"fmt"
	"unsafe"
)

//字符串类型
func main(){
	//基本使用
	var address string = "北京长城 beijingchangcheng"
	fmt.Printf("address=%s,address占用空间=%d,类型%T\n",address,unsafe.Sizeof(address),address)

	//反引号
	var str1 string = 
	`
	func main(){
		
	}
	`
	str2 := "abc\nabc"
	fmt.Printf("str1=%s,\nstr2=%s\n",str1,str2)

	//字符串拼接+
	str3 := "abc" 
	str4 :="def"
	str5 := str3 + str4
	fmt.Println("str5:",str5)

	//当一行太长时,需要使用多行字符串,要在最后用+,将+保留在上一行
	str6 := "abcdef"+
			"ghjklmm"
	fmt.Println("str6:",str6)
}

//字符串使用细节:
//1:Go统一使用UTF-8编码标识Unicode文本
//2:字符串一旦赋值就不能修改,在Go中字符串是不可变的
//3:字符串有两种表示形式:
//双引号,会识别转义字符
//反引号,以字符串的原生形式输出,包括换行和特殊字符,可以实现防止攻击,输出源码等效果
//4:字符串拼接+
//5:当一行太长时,需要使用多行字符串,要在最后用+