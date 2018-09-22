package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie 定義一部電影的資料,只有大寫開頭的欄位名稱才可以被匯出,如此才可以被json套件讀取欄位資料
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` //omitempty:若該欄位為false則不輸出
	Actors []string
}

func main() {
	movies := []Movie{
		{"Casablanca", 1942, false, []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{"Cool Hand Luke", 1967, true, []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	//indent縮排,第二參數：前綴字串,第三參數：縮排字串
	data1, err := json.MarshalIndent(movies, "", "--")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data1)

}
