# 思考步驟
1. 先嘗試使用github的api
2. 使用curl來get回應,同時使用-o參數,將回應存檔紀錄data.txt
   >`curl -o data.txt https://api.github.com/search/repositories?q=tetris+language:assembly&sort=stars&order=desc`
3. 觀察json結構，撰寫適合golang的struct
4. 