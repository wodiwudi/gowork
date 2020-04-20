package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"test/filelistingserver/filelist"
)

//定义一个type为http函数
type appHandler func(writer http.ResponseWriter, request *http.Request) error

//定义一个用户可见的错误类型
type userError interface {
	error
	Message() string
}

//errWrapper是统一错误处理函数参数为func(ResponsWriter,*Request)返回也是,把输入的函数包装，之后再输出
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		defer func() {
			res := recover()
			if res != nil {
				log.Error("panic : %v", err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		//type assertion
		if err != nil {
			if userError, ok := err.(userError); ok { //如果这个错误实现了用户级别错误则显示给用户浏览器
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return //不走下面的逻辑了
			}
			log.Warn("Error occured:%s", err.Error()) //记录日志，完整的错误信息
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(http.StatusNotFound), code) //外部报错应该输出普通信息
		}
	}
}

//文件服务器
func main() {
	http.HandleFunc("/", errWrapper(filelist.HandlerWeb))
	err := http.ListenAndServe(":8989", nil) //开启http服务 端口8989
	if err != nil {
		panic(err)
	}
}
