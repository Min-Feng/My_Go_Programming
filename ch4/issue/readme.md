# 思考步驟及筆記
1. 先嘗試使用github的api
   > https://developer.github.com/v3/search/
2. 使用curl來get回應,同時使用-o參數,將回應存檔紀錄data.txt
   >`curl -o data.txt https://api.github.com/search/repositories?q=tetris+language:assembly&sort=stars&order=desc`

    >api的多個參數,用&連接,如下有q,sort,order \
    q=tetris+language:assembly&sort=stars&order=desc

    >參數的值也可以用空白符號間隔開來,但記得要經過轉義QueryEscape \
    repo:golang/go is:open json decoder \
    空白的url編碼為 %20 \
    `https://api.github.com/search/issues?q=repo:golang/go%20is:open%20json%20decoder`

3. 觀察回傳的json結構，撰寫適合golang的struct
4. get请求傳遞参数的大小是有限制的，最大1024 bytes
5. 學習封裝package 將特定功能給其他程式使用 \
   要放在不同的資料夾 不然會有衝突
6. 考慮多種錯誤情況 給予適合的回應
7. http到回應後，處理錯誤情況，記得都要考慮resp.Body.Close()

   