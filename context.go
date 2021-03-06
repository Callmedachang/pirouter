package pirouter

import "net/http"

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	Path   string
	Method string
	Params map[string]string //TODO: params encode
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) NotFound() {
	c.Writer.WriteHeader(404)
	c.Writer.Write([]byte("NOT FOUND"))
}
