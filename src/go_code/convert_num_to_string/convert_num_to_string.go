package main

import(
	"fmt"
	"strconv"
)

//基本数据类型转string
func main(){
	//方式1:fmt.Sprintf("%参数",表达式),这个比较灵活,推荐
	//参数需要和表达式的数据类型相匹配
	//fmt.Sprintf()...会返回转换后的字符串
	var num1 int = 100
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string

	//需要掌握各种%转化的使用
	//https://studygolang.com/pkgdoc
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%v\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%v\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%v\n",str,str)

	str = fmt.Sprintf("%c",myChar)
	fmt.Printf("str type %T str=%v\n",str,str)

	//方式2:使用strconv包的参数
	var num3 int =99
	var num4 float64 = 543.3211
	var b2 bool = false

	//func FormatInt(i int64, base int) string
	//返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
	str = strconv.FormatInt(int64(num3),10)
	fmt.Printf("str type %T str=%v\n",str,str)

	//func Itoa(i int) string 很好用的函数
	str = strconv.Itoa(num3)
	fmt.Printf("str type %T str=%v\n",str,str)


	//func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	//函数将浮点数表示为字符串并返回。
	//bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	//fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、
	//             'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、
	//             'G'（指数很大时用'E'格式，否则'f'格式）。
	//prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。
	//                             如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	str = strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type %T str=%v\n",str,str)

	//func FormatBool(b bool) string
	//根据b的值返回"true"或"false"。
	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%v\n",str,str)


}