package main

import (
	"fmt"
)

// HasPrefix 確認前綴字母
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func main() {
	s := "sdf1243"
	comfirmPrefix := HasPrefix(s, "as")
	fmt.Println(comfirmPrefix)
}
