package file

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func openFile() *bufio.Scanner {
	file, err := os.Open("./players.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	return scanner
}

func GetRawList() []string {
	var rawList []string
	scanner := openFile()
	for scanner.Scan() {
		rawList = append(rawList, strings.ToLower(scanner.Text()))
	}
	return rawList
}
