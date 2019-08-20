package main

import (
	"fmt"

	"./file"
	"./list"
)

func main() {
	fmt.Println("P.T.S.S.")
	rawList := file.GetRawList()
	team1, team2 := list.GetTeams(rawList)
	fmt.Println(team1)
	fmt.Println(team2)

}
