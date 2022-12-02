package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// http://127.0.0.1:8001/B
	http.HandleFunc("/B", myHandler)
	http.ListenAndServe("127.0.0.1:8001", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "connect success!")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	fmt.Println("user name: ", r.URL.Query().Get("username"))

	// GET A
	resp, err := http.Get("http://127.0.0.1:8000/A")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("resp status:", resp.Status)
	fmt.Println("resp header:", resp.Header)
	fmt.Println("resp body:", resp.Body)

	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			w.Write([]byte(res))
			break
		}
	}

	if r.Method == "POST" {
		var (
			username string = r.PostFormValue("username")
			// password string = r.PostFormValue("password")
		)
		w.Write([]byte("username: " + username))
	}
	// w.Write([]byte(res))
	// w.Write([]byte(r.URL.Query().Get("username")))
	w.Write([]byte("\nB write back!\n"))
}
