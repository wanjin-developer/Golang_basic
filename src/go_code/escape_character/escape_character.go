package main

import "fmt"

//演示转义字符
func main() {
	// \t->tab
	fmt.Println("tom\tjack")
	// \\->\
	fmt.Println("C:\\sss")
	// /"->"
	fmt.Println("tom say \"sss\"")
	// /r->
	fmt.Println("tom say\rsss")
	// \n->Enter
	fmt.Println("姓名\t年龄\t籍贯\t住址\njoon\t12\t北京\t上海")
	//换行
	fmt.Println("asldkjasdkjlskjalkjdlkasjasldkjasdkjlskjalkj",
		"dlkasjasldkjasdkjlskjalkjdlkasjasldkjasdkjlskjalkjdlkasja",
		"sldkjasdkjlskjalkjdlkasjasldkjasdkjlskjalkjdlkasj")
}
