package handler

import (
	"fmt"
	"net/http"
)

type UserController struct {
}

func (uc UserController) ProcessRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Processing user request...")

}
