package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)

	//從標準輸入獲得資料,可利用重導向,將檔案的資料匯入
	//windows10 	"Get-Content .\testData.txt | go run .\dedup.go"
	//ubuntu18.04 	"go run .\dedup.go < testData.txt "
	//ref https://stackoverflow.com/questions/28504847/powershells-pipe-adds-linefeed/32981504#32981504
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "example: %v\n", err)
		os.Exit(1)
	}

}
