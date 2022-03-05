package main

import "fmt"

func main() {
	age := 50
	if age > 10 {
		fmt.Println("age is greater than 10")
	} else {
		fmt.Println("age is less than 10")
	}

	score := 80
	var grade string
	if score <= 50 {
		grade = "E"
	} else if score >= 60 {
		grade = "A"
	} else {
		grade = "C"
	}

	fmt.Println(grade)

	nilai := 3
	switch nilai {
	case 1:
		fmt.Println("nilai 1")
	case 2:
		fmt.Println("nilai 2")
	case 3:
		fmt.Println("nilai 3")
	default:
		fmt.Println("nilai tidak diketahui")
	}

	// perulangan for
	for i := 1; i <= 10; i++ {
		fmt.Println("sakura endo cantik", i)
	}

	// perulangan while
	i := 1
	for i <= 10 {
		fmt.Println("Aruno Nakanishi cantik", i)
		i++
	}
	a := 10
	for a <= 100 {
		fmt.Println("jepang", a)
		a++
	}

	title := "sakura miyawaki adalah idol jepang"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index", index, "letter", string(letter))
		}

	}

	for index, letter := range title {
		letterString := string(letter)
		// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		// 	fmt.Println("huruf", letterString, "ada di index", index)
		// }

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("huruf", letterString, "ada di index", index)

		}
	}

	var nama_idol [4]string
	nama_idol[0] = "sakura miyawaki"
	nama_idol[1] = "aruno nakanishi"
	nama_idol[2] = "sakura endo"
	nama_idol[3] = "Itou Miku"

	fmt.Println(nama_idol)

	var idol_grub = [4]string{"JKT48", "HKT48", "NOGIZAKA46", "AKB48"}
	fmt.Println(idol_grub)

	var vtuber = []string{"korone", "pekora", "moona"}
	// tanpa memberikan berapa jumlah array
	var newVtuber = vtuber[1:3]
	fmt.Println(newVtuber)
}
