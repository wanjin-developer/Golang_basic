package main

import (
	"fmt"
)

//定义一个变量保存华氏摄氏度,华氏温度转换成摄氏度的公式:5/9*(华氏温度-100),
//请求出华氏温度对应的华氏摄氏度
func main() {
	var wendu float32 = 134.2
	var sheshidu float32 = 5.0 / 9 * (wendu - 100)
	//华氏温度:134.2,华氏摄氏度:19
	fmt.Printf("华氏温度:%v,华氏摄氏度:%v", wendu, sheshidu)

}
