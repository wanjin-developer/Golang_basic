package main

import (
	"fmt"
	//为了使用utils.go文件的变量或者函数,我们需要先引入该包
	//go总是先从GOROOT出先搜索，再从GOPATH列出的路径顺序中搜索，
	//只要一搜索到合适的包就理解停止。当搜索完了仍搜索不到包时，将报错.
	//如果找不到包需要在环境变量中编辑GOPATH,添加该工程的路径
	"go_code/identifier/demo"
)

func main(){
	//报错:cannot refer to unexported name demo.heroName
	//fmt.Println(demo.heroName)

	fmt.Println(demo.HeroName)

}