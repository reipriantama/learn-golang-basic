package main

import "fmt"

func ViewTeam() {

}

func main() {
	var teamName, country, managerName string

	fmt.Println("Enter team name:")
	fmt.Scanln(&teamName)
	fmt.Println("Enter country:")
	fmt.Scanln(&country)
	fmt.Println("Enter manager name:")
	fmt.Scanln(&managerName)
}
