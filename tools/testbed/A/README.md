# http://127.0.0.1:8000/A
	fmt.Println(r.RemoteAddr, "connect success!")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)

    w.Write([]byte("A write back!"))

# ABCD
访问D，D向B和C发送post请求，B再向A发送get请求。
D->post B->get A
D->post C
postform内容如下：
{"username": {"usr123"}, "password": {"pwd123"}
