package main

import (
	"fmt"
	"os"
)

func main(){
	// 第一个参数是当前文件的路径
	var s, sep string
	for i :=0; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}