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
	fmt.Println(result.TotalCount)

	for i := -0; i < 5; i++ {
		fmt.Printf("%s %s %q\n", result.Items[i].Title, result.Items[i].User.Login, result.Items[i].UpdateTime)
	}
}
