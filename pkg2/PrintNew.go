package pkg2

import "fmt"

var v1 int = 5
var V1 int = 5

func Printnew1() {
	printnew("I am called from PrintNew.go")
}

func printnew(s1 string) {
	fmt.Println(s1)
}
