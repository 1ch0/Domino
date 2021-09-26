/**
  @Author NEO
  @Date 11:59 2021/9/18
  @Description //TODO
 **/

package main

import "net/http"
import "hades/framework"

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr: ':8080',
	}
	server.ListenAndServe()

	//fs := http.FileServer(http.Dir("/home/bob/static"))
	//http.Handle("/static/", http.StripPrefix("/static", fs))
}
