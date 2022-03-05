package main

import "fmt"

func main() {
	//map
	var idol map[string]string
	idol = make(map[string]string)

	idol["name"] = "sakura"
	idol["age"] = "20"
	idol["address"] = "jepang"
	fmt.Println(idol)

	idolgrub := map[string]string{
		"JKT48":     "jakarta",
		"AKIHABARA": "AKB48",
		"HAKATA":    "HKT48",
		"NAMBA":     "NMB48",
	}
	for key, val := range idolgrub {
		fmt.Println(key, " \t:", val)
	}

	// slice of map
	students := []map[string]string{
		{"name": "sakura", "age": "20", "address": "jepang"},
		{"name": "aruno", "age": "20", "address": "jepang"},
		{"name": "sakura endo", "age": "20", "address": "jepang"},
	}

	// fmt.Println(students)
	for _, student := range students {
		// fmt.Println(student)
		fmt.Println(student["name"], "umurnya", student["age"], "asal dari ", student["address"])
	}

	// kuis
	scores := [8]int{80, 90, 70, 60, 50, 40, 30, 20}

	var total int

	for _, score := range scores {
		total = total + score
	}

	length := len(scores)
	average := float64(total) / float64(length)
	fmt.Println(average)

	// function dasar
	printMyResult("ini function dengan parameter")
	text := printWithReturn("ini function dengan return")
	fmt.Println(text)

	penjumlahan := add(10, 20)
	fmt.Println(penjumlahan)

}

// basic funtion

func printMyResult(sentence string) {
	fmt.Println(sentence)
}

func printWithReturn(sentence string) string {
	newSentence := sentence + " ini adalah function dengan return"
	return newSentence
}

func add(a, b int) int {
	return a + b
}
