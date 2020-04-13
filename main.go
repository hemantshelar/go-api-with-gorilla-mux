package main

import (
	"github.com/hshelar/api-structure-demo/app"
)

func main() {
	a := app.App{}
	a.Init()
	//a.Run(":3000")
	a.RunWithGorila(":3000")
}
