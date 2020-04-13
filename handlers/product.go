package handler

import "net/http"

type ProductController struct{}

func (this ProductController) ProcessRequest(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Serving products..."))
}
