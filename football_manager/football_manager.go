package main

import (
	"fmt"
)

// Struct untuk player
type Player struct {
	Name     string
	Age      int
	Position string
}

// Struct untuk team
type Team struct {
	Name        string
	Country     string
	ManagerName string
	Players     []Player
}

// Menambahkan player ke dalam team
func (t *Team) AddPlayer(player Player) {
	t.Players = append(t.Players, player)
	fmt.Println("Player added successfully.")
}

// Update player informasi di team
func (t *Team) UpdatePlayer(player Player) {
	for i, p := range t.Players {
		if p.Name == player.Name {
			t.Players[i] = player
			fmt.Println("Player updated successfully.")
			return
		}
	}
	fmt.Println("Player not found.")
}

// menghapus player dari team
func (t *Team) DeletePlayer(name string) {
	for i, p := range t.Players {
		if p.Name == name {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			fmt.Println("Player deleted successfully.")
			return
		}
	}
	fmt.Println("Player not found.")
}

// menampilkan informasi terkait player dan team
func (t *Team) ViewTeam() {
	fmt.Printf("Team Name: %s\n", t.Name)
	fmt.Printf("Country: %s\n", t.Country)
	fmt.Printf("Manager: %s\n", t.ManagerName)
	fmt.Println("Players:")
	for _, player := range t.Players {
		fmt.Printf("Name: %s, Age: %d, Position: %s\n", player.Name, player.Age, player.Position)
	}
}

func main() {
	var teamName, country, managerName string

	// Input team name, country, dan manager name
	fmt.Println("Enter team name:")
	fmt.Scanln(&teamName)
	fmt.Println("Enter country:")
	fmt.Scanln(&country)
	fmt.Println("Enter manager name:")
	fmt.Scanln(&managerName)

	// Membuat team
	team := Team{Name: teamName, Country: country, ManagerName: managerName}

	// Add players to the team
	fmt.Println("Add players to the team (press Enter after each player, type 'done' when finished):")
	for {
		var playerName, position string
		var age int
		fmt.Println("Player name:")
		fmt.Scanln(&playerName)
		if playerName == "done" {
			break
		}
		fmt.Println("Age:")
		fmt.Scanln(&age)
		fmt.Println("Position:")
		fmt.Scanln(&position)
		player := Player{Name: playerName, Age: age, Position: position}
		team.AddPlayer(player)
	}

	for {
		var choice string
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Add player")
		fmt.Println("2. Update player")
		fmt.Println("3. Delete player")
		fmt.Println("4. View team")
		fmt.Println("5. Exit")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var playerName, position string
			var age int
			fmt.Println("\nAdd a new player:")
			fmt.Println("Player name:")
			fmt.Scanln(&playerName)
			fmt.Println("Age:")
			fmt.Scanln(&age)
			fmt.Println("Position:")
			fmt.Scanln(&position)
			player := Player{Name: playerName, Age: age, Position: position}
			team.AddPlayer(player)
		case "2":
			var updatedPlayer Player
			fmt.Println("\nEnter player name to update:")
			fmt.Println("Player name:")
			fmt.Scanln(&updatedPlayer.Name)
			fmt.Println("Age:")
			fmt.Scanln(&updatedPlayer.Age)
			fmt.Println("Position:")
			fmt.Scanln(&updatedPlayer.Position)
			team.UpdatePlayer(updatedPlayer)
		case "3":
			var deletePlayerName string
			fmt.Println("\nEnter player name to delete:")
			fmt.Scanln(&deletePlayerName)
			team.DeletePlayer(deletePlayerName)
		case "4":
			fmt.Println("\nTeam Information:")
			team.ViewTeam()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
