package main

import "fmt"

func main() {
	var nomorA int = 48
	var nomorB *int = &nomorA

	fmt.Println("address dari nomor a pada nomorb ", nomorB)
	fmt.Println("nilai dari nomor a pada nomorb ", *nomorB)

	fmt.Println("address dari nomor a pada nomorb ", &nomorA)
	fmt.Println("nilai dari nomor a pada nomorb ", nomorA)
}
