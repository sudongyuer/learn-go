package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("开始处理异常")
		// 获取异常信息
		if err := recover(); err != nil {
			//  输出异常信息
			fmt.Println("error:", err)
		}
		fmt.Println("结束异常处理")
	}()
	generateError()
}

func generateError() {
	i := 10 / 1
	fmt.Println(i)
	panic("异常信息")
}
