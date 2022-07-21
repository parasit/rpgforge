package rpgforge

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed names/all.last
var fLastName embed.FS

//go:embed names/male.first
var fFirstMaleName embed.FS

//go:embed names/female.first
var fFirstFemaleName embed.FS

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Read data from embed file and return []string
func ReadEmbedData(embedFile* embed.FS, filename string) []string {
	data, err := embedFile.ReadFile(filename)
	check(err)
	myList := strings.Split(string(data[:]), "\n")
	return myList
}

// Get random last name
func GetLastName() string {
	var result string
	initialize()
  myList := ReadEmbedData(&fLastName, "names/all.last")
	fmt.Printf("Got %d last names\n", len(myList))
	result = strings.Title(strings.ToLower(myList[rand.Intn(len(myList))]))
	return result
}
func GetFirstMaleName() string {
	var result string
	initialize()
  myList := ReadEmbedData(&fFirstMaleName, "names/male.first")
	fmt.Printf("Got %d first male names\n", len(myList))
	result = strings.Title(strings.ToLower(myList[rand.Intn(len(myList))]))
	return result
}
func GetFirstFemaleName() string {
	var result string
	initialize()
  myList := ReadEmbedData(&fFirstMaleName, "names/female.first")
	fmt.Printf("Got %d first female names\n", len(myList))
	result = strings.Title(strings.ToLower(myList[rand.Intn(len(myList))]))
	return result
}
