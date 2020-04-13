package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Route
	Db     string
}

func (app *App) Init() {
	fmt.Println("Applicatin init...")
}

func (app *App) Run(port string) {
	fmt.Println("Starting webserver...")
	http.HandleFunc("/", func(r http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL.Path)
		fmt.Println(request.Method)
	})
	http.ListenAndServe(port, nil)
}

func (app *App) RunWithGorila(port string) {
	fmt.Println("Starting webserver with Gorila...")
	r := mux.NewRouter()
	r.HandleFunc("/", func(r http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL.Path)
		fmt.Println(request.Method)
	})
	http.ListenAndServe(port, r)
}
