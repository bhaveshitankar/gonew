package pkg2

import "fmt"

func Printnew2() {
	printnew(fmt.Sprintf("abcd%s", "1234"))
	fmt.Printf("abcd%s", "1234")
}
