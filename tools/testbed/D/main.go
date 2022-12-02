package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// http://127.0.0.1:8003/D
	http.HandleFunc("/D", myHandler)
	http.ListenAndServe("127.0.0.1:8003", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "connect success!")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)

	// msg := make(map[string]string)
	// msg["username"] = "usr123"
	// msg["password"] = "pwd123"
	// resqMsg, _ := json.Marshal(msg)

	// POST B
	// resp, err := http.NewRequest("POST", "http://127.0.0.1:8001/B", reqUser)
	// resp, err := http.Post("http://127.0.0.1:8001/B", "application/json;charset=utf-8", bytes.NewBuffer([]byte(resqMsg)))
	resp, err := http.PostForm("http://127.0.0.1:8001/B", url.Values{"username": {"usr123"}, "password": {"pwd123"}})
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("resp url:", resp.Status)
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

	// POST C
	// respC, err := http.NewRequest("POST", "http://127.0.0.1:8002/C", reqUser)
	//respC, err := http.Post("http://127.0.0.1:8002/C", "application/json;charset=utf-8", bytes.NewBuffer([]byte(resqMsg)))

	respC, err := http.PostForm("http://127.0.0.1:8002/C", url.Values{"username": {"usr123"}, "password": {"pwd123"}})
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer respC.Body.Close()
	fmt.Println(respC.Status)
	fmt.Println(respC.Request.URL)
	fmt.Println(respC.Header)
	fmt.Println(respC.Body)

	bufC := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := respC.Body.Read(bufC)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(bufC[:n])
			fmt.Println(res)
			w.Write([]byte(res))
			break
		}
	}

	// w.Write([]byte(res))
	w.Write([]byte("\nD write back!\n"))
}
