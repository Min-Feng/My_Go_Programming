package main

import "crypto/sha256"
import "fmt"

func main() {
	xSlice := []byte("x")

	/*
		.Sum256回傳 checksum的值 共32bytes

		想得到字串checksum的值有兩種方法

		第一种调用方法
		sum := sha256.Sum256([]byte("hello world\n"))
		fmt.Printf("%x\n", sum)

		第二种调用方法
		h := sha256.New()
		h.Write([]byte("hello world\n"))
		fmt.Printf("%x\n", h.Sum(nil))

		ref : https://blog.csdn.net/luckydog612/article/details/80547758
			  https://go-zh.org/pkg/crypto/sha256/#Sum256
			  https://go-zh.org/src/crypto/sha256/sha256.go?s=1643:1663#L72
	*/
	c1 := sha256.Sum256(xSlice)

	/*
		%s   输出字符串表示（string类型或[]byte)
		%q   双引号围绕的字符串
		%x   十六进制，小写字母，每字节两个字符
		%T   顯示變數的型態
	*/
	fmt.Printf("%q\n%x\n%T\n", xSlice, c1, c1)
}
