package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //解析form表单内容,不写不会解析
	if err != nil {
		log.Fatal("Form: ", err)
	}
	fmt.Println(r.Form)               //这些信息是输出到服务器端的打印信息
	fmt.Println("path :", r.URL.Path) //url路径信息
	fmt.Println("SCHEME :", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "hello astaxie") //写入w的事输出给客户端的
}

//登录handler
func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //解析form表单内容,不写不会解析
	if err != nil {
		log.Fatal("Form: ", err)
	}
	fmt.Println("method:", r.Method)
	//根据方法来进行逻辑的处理
	if r.Method == "GET" {
		strNowTime := time.Now().Unix()
		fmt.Printf("%T %v\n", strNowTime, strNowTime)
		h := md5.New()                                                 //实例化一个md5
		_, err := io.WriteString(h, strconv.FormatInt(strNowTime, 10)) //向客户端输出
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(token)                                    //589991eb1b7e13b6373f5ceea45f35a8
		fmt.Printf("%T", strconv.FormatInt(strNowTime, 10))   //将int64转换string
		t, err := template.ParseFiles("src/goweb/login.html") //解析模板文件
		if err != nil {
			log.Fatal("Form: ", err)
		}
		e := t.Execute(w, token) //将数据渲染到已经解析的模板
		if e != nil {
			log.Fatal("Form: ", err)
		}
	} else {
		_, _ = strconv.Atoi(r.Form.Get("age")) //返回10进制的strconv.ParseInt()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		//预防xss跨站脚本攻击
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端

	}

}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		times := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(times, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(token)
		t, _ := template.ParseFiles("src/goweb/update.html")
		err := t.Execute(w, token)
		if err != nil {
			log.Fatal("err:", err)
		}

	} else {
		token := r.Form
		fmt.Println("111", token)
		if token != nil {
			fmt.Println(token)
			err := r.ParseMultipartForm(32 << 20)
			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				log.Fatal("upload", err)
			}
			defer file.Close()
			_, err = fmt.Fprintf(w, "%v", handler.Header)
			f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 066)
			defer f.Close()
			io.Copy(f, file)
		} else {

		}
	}
}

func main() {
	//用go build 再运行.exe也可以开启服务
	//http.HandleFunc("/", sayhelloName) //设置访问路由为根目录，接收方法为sayhelloName(控制器)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
