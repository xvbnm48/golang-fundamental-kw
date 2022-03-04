package main

import (
	"fmt"

	"github.com/xvbnm48/golang-fundamental-kw/part1/calculation"
)

func main() {
	fmt.Println("Hello, World!")
	result := calculation.Add(10, 20)
	//sentence := TestAja()
	//fmt.Println(sentence)
	fmt.Println(result)

	fmt.Println("proses perkalian :")
	result2 := calculation.ProsesKali(10, 10)
	fmt.Println(result2)

	var nama string = "sakura miyawaki"
	fmt.Println(nama)
}
