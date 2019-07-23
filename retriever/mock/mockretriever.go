package mock

import "fmt"

type Retriever struct {
	Contents string
}

// 实现内置的 Stringer 接口
func (r *Retriever) String() string {
	return fmt.Sprintf("Retiever: {Contents=%s}", r.Contents)
}

// 实现 Poster 接口
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

// 实现 Retriever 接口
func (r *Retriever) Get(url string) string {
	return r.Contents
}

