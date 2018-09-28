package main

import (
	"fmt"
	"log"
	"myGoPractice/ch4/issue/github" //以資料夾為單位 跟主程式要放在不同資料夾
	"os"
)

func main() {
	para := os.Args[1:]
	result, err := github.SearchIssues(para)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%[1]d %[2]s\n", result.TotalCount, "issues")

	for i := 1; i < 10; i++ {
		fmt.Printf("%20s %v %s\n", result.Items[i].User.Login, result.Items[i].UpdateTime, result.Items[i].Title)
	}
}
