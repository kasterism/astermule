# http://127.0.0.1:8001/B

	fmt.Println(r.RemoteAddr, "connect success!")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	fmt.Println("user name: ", r.URL.Query().Get("username"))

	// GET A
	resp, err := http.Get("http://127.0.0.1:8000/A")

	w.Write([]byte(res))
	w.Write([]byte("username: " + username))
	w.Write([]byte("\nB write back!\n"))

# ABCD
访问D，D向B和C发送post请求，B再向A发送get请求。
D->post B->get A
D->post C
postform内容如下：
{"username": {"usr123"}, "password": {"pwd123"}

