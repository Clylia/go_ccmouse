package _9_文件服务器

/**
|-------------------------------------------------------------------------------------------------
|
|	实现一个简易的 web 服务器
|
|-------------------------------------------------------------------------------------------------
*/

import (
	"github.com/gpmgo/gopm/log"
	"learn_go_ccmouse/19_文件服务器/filelisting"
	"net/http"
	_ "net/http/pprof"
	"os"
)

// 定义一个错误类型函数
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 对错误信息做详细处理
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// panic 处理，错误捞出处理
		defer func() {
			if r := recover(); r != nil {
				log.Warn("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		// 对错误信息进行判断并输出
		if err != nil {
			// log 错误信息
			log.Warn("Error handling request: %s", err.Error())

			// 处理 user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			// 处理系统错误 system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	// 定义一个请求规则
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	// 开启一个端口监听
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
