//一個資料夾內，一定要有main package ,也只有一個main 函式
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], "-*-"))
}
