package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http://127.0.0.1:8000/A
	http.HandleFunc("/A", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "connect success!")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)

	w.Write([]byte("\nA write back!\n"))
}
