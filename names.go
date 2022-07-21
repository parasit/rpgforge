package rpgforge

import (
	"embed"
	"math/rand"
	"strings"

	"golang.org/x/text/cases"
  "golang.org/x/text/language"
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

// Read data from embed file and return string
func GetRandomEmbedData(embedFile* embed.FS, filename string) string {
	data, err := embedFile.ReadFile(filename)
	check(err)
	myList := strings.Split(string(data[:]), "\n")
  c := cases.Title(language.Und)
  result := c.String(strings.ToLower(myList[rand.Intn(len(myList))]))
	return result
}

// Get random last name
func GetLastName() string {
	var result string
	initialize()
  result = GetRandomEmbedData(&fLastName, "names/all.last")
	// result = strings.Title(strings.ToLower(myList[rand.Intn(len(myList))]))
	return result
}
func GetFirstMaleName() string {
	var result string
	initialize()
  result = GetRandomEmbedData(&fFirstMaleName, "names/male.first")
	return result
}
func GetFirstFemaleName() string {
	var result string
	initialize()
  result = GetRandomEmbedData(&fFirstFemaleName, "names/female.first")
	return result
}
