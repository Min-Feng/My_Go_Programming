package main

import (
	"log"
	"net/http"
	"fmt"
	)

func main(){
	http.HandleFunc("/",hander)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func hander(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"URL.Path=%s",r.URL.Path)
}