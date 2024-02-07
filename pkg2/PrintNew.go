package pkg2

import "fmt"

func Printnew1() {
	printnew("I am called from PrintNew.go")
}

func printnew(s1 string) {
	fmt.Println(s1)
}
