package utils

import "fmt"

var Age int
var Name string

func init() {
	Age = 20
	Name = "tom"
	fmt.Println("utils-init")
}
