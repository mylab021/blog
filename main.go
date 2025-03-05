package main

import (
	"fmt"
	"github.com/mylab021/msgo"
	"net/http"
)

func main() {

	// Go的原生http路由
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "%s 欢迎来到码神之路GoWeb教程", "Chris.Zhang")
	// })

	//
	engine := msgo.New()
	g := engine.Group("user")
	g.Add("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s 欢迎来到码神之路GoWeb教程", "Chris.Zhang")
	})
	engine.Run()

}
