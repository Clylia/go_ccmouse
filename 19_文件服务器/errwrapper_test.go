package _9_文件服务器

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// panic 类型错误测试
func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

// user error 类型错误测试
type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

// 404 Not Found
func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

// 403 Forbidden
func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

// Unknown
func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unknown error")
}

// noError
func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// 测试 errWrapper()
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()	// 创建一个假的 response
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)	// 创建一个假的 request
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

// 测试整体的这个 web 服务器
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))	// 创建一个真的request
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

// 验证返回值
func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T)  {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect (%d, %s); got (%d, %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}
