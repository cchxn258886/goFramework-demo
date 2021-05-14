package main

import (
	"context"
	"net/http"
)

//测试
func main() {
	//done := make(chan error,2);
	//stop := make(chan struct{})
	//go func() {
	//	done <- server()
	//}()
	var s = ""
	println(len(s))
	if s == "" {
		println("true")
	}
}
func server(addr string, handle http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handle,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}
