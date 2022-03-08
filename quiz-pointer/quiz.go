package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (game *Gamer) AddGame(nama string) {
	game.Games = append(game.Games, nama)
}

func main() {
	gamer := Gamer{Name: "Zelda"}

	gamer.AddGame("Super Mario")
	gamer.AddGame("Street Fighter")
	gamer.AddGame("GTA")
	gamer.AddGame("Mario Kart")
	gamer.AddGame("mobile legends")
	gamer.AddGame("Call of Duty")
	for _, game := range gamer.Games {
		fmt.Println(game)
	}

}
