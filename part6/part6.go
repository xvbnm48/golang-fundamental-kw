package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " graduated"
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}

	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error eccured", r)
	} else {
		fmt.Println("application running normally")
	}
}

func main() {
	defer catch()
	student := Student{1, "sakura endo", 3.5}
	fmt.Println(student)

	student.graduate()
	fmt.Println(student.Name)

	// error, panic, recover

	var input string
	fmt.Println("Enter some number: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, " is number")
	} else {
		fmt.Println(input, " is not number")
		fmt.Println(err.Error())
	}
	// =================================

	// panic
	var anime string
	fmt.Println("enter your favorite anime ")
	fmt.Scanln(&anime)
	if valid, err := validate(anime); valid {
		fmt.Println("favorite anime is ", anime)
	} else {
		panic(err.Error())
	}

	// name

	var name string
	fmt.Println("enter your name ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("hallo ", name)
	} else {
		fmt.Println(err.Error())
	}

}
