package list

import (
	"fmt"
	"strings"
)

func GetTeams(rawList []string) (team1 []string, team2 []string) {
	team1, team2 = getGk(rawList)
	zgList := getAllPositions(rawList, "zag")
	alaList := getAllPositions(rawList, "ala")
	meiaList := getAllPositions(rawList, "meia")
	ataList := getAllPositions(rawList, "ata")

	fmt.Println(zgList)
	fmt.Println(alaList)
	fmt.Println(meiaList)
	fmt.Println(ataList)

	// fmt.Println(rawList)

	return team1, team2
}

func getGk(rawList []string) (team1 []string, team2 []string) {
	keyWord := "gk"
	team1 = append(team1, iterateGetPosition(rawList, keyWord))
	team2 = append(team2, iterateGetPosition(rawList, keyWord))
	return team1, team2
}

func iterateGetPosition(rawList []string, keyWord string) string {
	for i, player := range rawList {
		if strings.Contains(player, keyWord) {
			position := strings.Split(player, "-")[1]
			(rawList)[i] = ""
			return position
		}
	}
	return ""
}

func getAllPositions(rawList []string, position string) []string {
	var tempList []string
	for _, player := range rawList {
		if strings.Contains(player, position) {
			tempList = append(tempList, player)
		}
	}
	return tempList
}
