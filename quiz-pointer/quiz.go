package main

import (
	"errors"
	"fmt"
	"strings"
)

type Gamer struct {
	Name  string
	Games []string
}

func (game *Gamer) AddGame(nama string) {
	game.Games = append(game.Games, nama)
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error eccured", r)
	} else {
		fmt.Println("application running normally")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}

	return true, nil
}

func main() {
	// gamer := Gamer{Name: "Zelda"}

	// gamer.AddGame("Super Mario")
	// gamer.AddGame("Street Fighter")
	// gamer.AddGame("GTA")
	// gamer.AddGame("Mario Kart")
	// gamer.AddGame("mobile legends")
	// gamer.AddGame("Call of Duty")
	// for _, game := range gamer.Games {
	// 	fmt.Println(game)
	// }
	defer catch()
	var name string
	fmt.Println("enter your name ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("hallo ", name)
	} else {
		fmt.Println(err.Error())
	}
}
