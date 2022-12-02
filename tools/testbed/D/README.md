# http://127.0.0.1:8003/D
	
    fmt.Println(r.RemoteAddr, "connect success!")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)

    // POST B
	resp, err := http.PostForm("http://127.0.0.1:8001/B", url.Values{"username": {"usr123"}, "password": {"pwd123"}})
    w.Write([]byte(res))

    // POST C
    respC, err := http.PostForm("http://127.0.0.1:8002/C", url.Values{"username": {"usr123"}, "password": {"pwd123"}})
    w.Write([]byte(res))

    w.Write([]byte("D write back!\n"))

# ABCD
访问D，D向B和C发送post请求，B再向A发送get请求。
D->post B->get A
D->post C
postform内容如下：
{"username": {"usr123"}, "password": {"pwd123"}