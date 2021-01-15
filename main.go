package main

import (
	"log"
	"net/http"
	"os"
)

/*func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Service Started..") //这个写入到w的是输出到客户端的
	r.ParseForm()
	var String string
	//if r.Method == "GET" {
	//	String = r.FormValue("String")
	//} else if r.Method == "POST" {
	//	String = r.PostFormValue("String")
	//}
	//io.WriteString(w, "String is:"+String)
	//* GET & POST test
}*/
func main() {
	args := os.Args
	var PORT string
	PORT = args[1]
	//const PORT = "8090"
	http.Handle("/", http.FileServer(http.Dir("www"))) //设置访问的路由
	err := http.ListenAndServe(":"+PORT, nil)          //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
