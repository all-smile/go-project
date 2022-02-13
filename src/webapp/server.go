package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(write, "Hello world, %s!", request.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
