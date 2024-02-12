/*
1. go is faster
*/
// Hi

package main

import (
	"fmt"

	"github.com/bhaveshitankar/gonew/pkg1"
	pkg1pkg1 "github.com/bhaveshitankar/gonew/pkg1/pkg1"
)

func packageCaller() {
	pkg1.HelloPrinter()
	pkg1pkg1.HelloPrinter()
}

var m1 int

func dataTypes() {
	var a int = 10
	var a1 = 10
	var aSlice []int                // its a slice
	var aArr1 [2]int = [2]int{1, 2} // its a array
	var aArr11 = [2]int{1, 2}       // its a array
	arr2 := [...]int{1, 2}          // its a array
	aSlice12 := []int{}
	aArr12 := [5]int{}
	aSlice10 := make([]int, 10) //  we want the slice to have 10elements
	fmt.Println(a, a1, aSlice, aArr1, aArr11, arr2, aSlice12, aArr12, aSlice10)
	/*
		var b int = 20
		a1, b2 := 10, 20
		// define floats
		var c float64 = 67.54
		c1 := 3.14
		// define bool
		var a2 bool = true
		a3 := false
		a3 = true
		// define string, byte(8), rune(16)
		var hello string = "v1"
		hello = "v3"
		hello1 := "golang go"
		hello2 := `ŋ`
		for h := range hello2 {
			fmt.Printf("%T,%v\n", hello2[h], hello2[h])
		}
		// define a constant
		const const1 = 20
		const (
			v1 = 2
			v2 = "dgdgdg"
		)
		// const1 = 5 -- not possible

		// define a array 0,1,2,3,4,5
		var arr = [6]int{1, 2} // length
		arr[2], arr[3] = 3, 4
		fmt.Println(arr[3]) // 0
		// define a slice
		sliceZeros := arr[3:]                      // length, capacity
		sliceZeros = append(sliceZeros, arr[:]...) // 0,0,0   1,2,3,4,0,0
		// arr --> 1234 --> append(sliceZeros, 1,2,3,4)
		// cap = 4*32bits--> [32][32][32][32]  [32][32][32][32]

		fmt.Println(arr[0], arr[2], len(sliceZeros), cap(sliceZeros))*/
}

func main() {
	dataTypes()
}
