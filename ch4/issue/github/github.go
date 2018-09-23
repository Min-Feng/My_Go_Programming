package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url" //url.QueryEscape来对查询中的特殊字符进行转义操作
	"strings" //為了將api參數加入到url
	"time"
)

//GitHubIssueURL api網址
const GitHubIssueURL = "https://api.github.com/search/issues"

//SearchIssueResult 搜尋結果
type SearchIssueResult struct {
	TotalCount int `json:"total_count"`

	//會接收許多data,應該用slice型別接收，並用指標讀取，避免傳輸整個struct，效率差
	//Items      IssueItem
	Items []*IssueItem
}

//IssueItem 項目資料
type IssueItem struct {
	HTMLURL    string `json:"html_url"`
	Title      string
	User       *UserData //用指標讀取，避免傳輸整個struct，效率差
	UpdateTime time.Time `json:"updated_at"`
}

//UserData 使用者資料
type UserData struct {
	Login   string
	HTMLURL string `json:"htnl_url"`
}

//SearchIssueResult2 嘗試將所有結構 寫在一個struct
type SearchIssueResult2 struct {
	TotalCount int `json:"total_count"`
	Items      []*struct {
		HTMLURL string `json:"html_url"`
		Title   string
		User    struct {
			Login   string
			HTMLURL string `json:"htnl_url"`
		}
		UpdateTime time.Time `json:"updated_at"`
	}
}

//SearchIssues golang傳統，第一個回傳正常值，第二個回傳錯誤值
//為了不傳遞大量資料，通常回傳指標，用指標進行讀寫
//考慮多種錯誤情況 給予適合的回應
func SearchIssues(query []string) (*SearchIssueResult2, error) {
	q := url.QueryEscape(strings.Join(query, " "))    //為了轉義空白符號 使用QueryEscape
	resp, err := http.Get(GitHubIssueURL + "?q=" + q) //同為字串 可以用加號 串連
	if err != nil {
		return nil, err //解析失敗,回傳空的結果
	}

	//get要求得到回應,但要考慮到回應是正確,記得close開啟過的檔案
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	//解析json失敗的情況
	var result SearchIssueResult2
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { //&result 目前不懂原理，之後要搞懂
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
