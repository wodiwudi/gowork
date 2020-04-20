package filelist

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userErrorString string //定义一个string type 实现web.go的userError接口

func (s userErrorString) Error() string {
	return s.Message()
}

func (s userErrorString) Message() string {
	return string(s)
}

const prefix = "/list/"

func HandlerWeb(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userErrorString("path must start with" + prefix)
	}
	path := request.URL.Path[len(prefix):] //获取url路径
	fmt.Println(request.URL.Path)
	file, err := os.Open(path) //打开本地文件
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file) //读取文件所有内容
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
