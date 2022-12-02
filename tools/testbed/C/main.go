package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http://127.0.0.1:8002/C
	http.HandleFunc("/C", myHandler)
	http.ListenAndServe("127.0.0.1:8002", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "connect success!")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// fmt.Println("user name: ", r.URL.Query().Get("username"))

	if r.Method == "POST" {
		var (
			username string = r.PostFormValue("username")
			// password string = r.PostFormValue("password")
		)
		w.Write([]byte("username: " + username))
	}

	// w.Write([]byte("username: " + r.URL.Query().Get("username")))
	w.Write([]byte("\nC write back!\n"))
}
