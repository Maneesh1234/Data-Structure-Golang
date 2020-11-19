package main

// package main

import (
	"fmt"
)

// func main() {

// 	fmt.Printf("%U\n", []rune("0b£"))
// 	rPound := '£'
// 	fmt.Printf("Type: %s\n", reflect.TypeOf(rPound))
// 	fmt.Printf("Unicode CodePoint: %U\n", rPound)
// 	//fmt.Printf("Character: %c\n", r)
// }

func main() {
	r := 'a'

	//Print Size
	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

	//Print Type
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	//Print Code Point
	fmt.Printf("Unicode CodePoint: %U\n", r)

	//Print Character
	fmt.Printf("Character: %c\n", r)
	s := "0b£"

	//This will print the Unicode Points
	fmt.Printf("%U\n", []rune(s))

	//This will the decimal value of Unicode Code Point
	fmt.Println([]rune(s))
}

Rune array to string
func main() {
	runeArray := []rune{'a', 'b', '£'}
	s := string(runeArray)
	fmt.Println(s)
}

// //String to Rune Array

// func main() {
// 	s := "ab£"
// 	r := []rune(s)
// 	fmt.Printf("%U\n", r)
// }
