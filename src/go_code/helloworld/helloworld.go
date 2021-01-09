//开发一个hello.go程序，输出“hello，world！”

//表示该hello.go文件所在的包是main，在go中，每个文件都必须归属于一个包
package main	

//表示引入一个包，包名fmt，引入该包后就可使用fmt函数
import "fmt"	

//func是一个关键字，表示一个函数；main是函数名，是一个主函数即是程序的入口
func main(){
	//表示调用fmt的函数Println输出“hello,world”
	fmt.Println("hello,world!")	
}

//控制台使用 go build hello.go 编译，或者使用go run 直接运行


//go执行流程：
//1、编译 go build 生成可执行文件
//2、运行

//go格式化命令：gofmt -w xxxx.go