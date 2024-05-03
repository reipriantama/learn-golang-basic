package main

import (
	"fmt"
)

type Player struct {
	Name     string
	Position string
	Age      int
}

type Team struct {
	Name    string
	Players []Player
}

func addPlayerToTeam(team *Team, player Player) {
	team.Players = append(team.Players, player)
}

func viewTeam(team Team) {
	fmt.Println("Team:", team.Name)
	fmt.Println("Players:")
	for _, player := range team.Players {
		fmt.Println("- Name:", player.Name, "Position:", player.Position, "Age:", player.Age)
	}
}

func main() {
	myTeam := Team{Name: "My Team"}

	var playerName, playerPosition string
	var playerAge int

	fmt.Println("Enter player's name:")
	fmt.Scanln(&playerName)

	fmt.Println("Enter player's position:")
	fmt.Scanln(&playerPosition)

	fmt.Println("Enter player's age:")
	fmt.Scanln(&playerAge)

	newPlayer := Player{Name: playerName, Position: playerPosition, Age: playerAge}
	addPlayerToTeam(&myTeam, newPlayer)

	viewTeam(myTeam)
}
