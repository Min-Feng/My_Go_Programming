package main

/*
flag套件提供了一系列解析命令行参数的功能
如果是寫大型 Web Application
不推薦使用 flag，原因是 flag 不支援讀取系統環境變數

命令行 flag 表示法有如下三种形式：
-flag 		// 只支持bool类型，也就是要使用flag.Bool
-flag=x     // 最穩健的表示法,不用管參數的型別
-flag x 	// 只支持非bool类型

ref: https://www.jianshu.com/p/f9cf46a4de0e
*/

import (
	"flag"
	"fmt"
	"strings"
)

//第一個參數：參數名稱 第二個參數：參數預設值 第三個參數：無效參數或-h的說明文
var n = flag.Bool("n", false, "設定為是否換行")
var s = flag.String("s", "**", "設定分隔符號")

func main() {
	//flag套件中，底層(underlying)使用了os.Arg,因此讀取命令列,不需要使用os.Arg
	//flag.Parse解析了命令列的各個參數，並且區分為flag參數 及 非flag參數
	flag.Parse()

	/*
		flag.Args為 非flag參數 的slice集合

		e.g.
		$ ./echo4 -s / a bc def
		-s 		： flag參數
		/ 		： flag參數的 值
		a bc def： 非flag參數 */
	fmt.Print(strings.Join(flag.Args(), *s))

	if !*n {
		fmt.Println()
	}

	/* bool值適用上述談到的 命令行flag 表示法三种形式的第一種

	如果在命令列 出現了-n  則表示n=true
	$ go run echo4.go -n
	true

	如果命令列 沒有出現-n  則n為預設值false
	$ go run echo4.go
	false
	*/
	fmt.Println(*n)
}
