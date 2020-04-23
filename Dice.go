package rpgforge

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// Dice Basic dice type
type Dice struct {
	Sides int
}

// DiceHand struct contains dice and number of throws
type DiceHand struct {
	dice   Dice
	throws int
}

// DiceBag contains results of throw
type DiceBag struct {
	result   []int
	modifier int
}

func initialize() {
	rand.Seed(time.Now().UnixNano())
}

// ThrowDices takes dice string (like 3d6) and parse it
func ThrowDices(diceString string) DiceBag {
	// var paramsMap map[string]string
	var throws, dice int
	var err error
	exploding := false
	var paramsMap map[string]string
	matchTest, _ := regexp.MatchString("^(?P<throws>\\d*)(?P<symbol>[dDkKfF])(?P<dice>\\d+)(?P<exploding>!*)(?P<modifier>[+-]*)(?P<modValue>\\d*)$", diceString)
	if !matchTest {
		db := DiceBag{make([]int, 0), 0}
		return db
	}
	reg := regexp.MustCompile("^(?P<throws>\\d*)(?P<symbol>[dDkKfF])(?P<dice>\\d+)(?P<exploding>!*)(?P<modifier>[+-]*)(?P<modValue>\\d*)$")
	match := reg.FindStringSubmatch(diceString)
	paramsMap = make(map[string]string)
	for i, name := range reg.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	// fmt.Println(paramsMap)
	if len(paramsMap["throws"]) > 0 {
		if throws, err = strconv.Atoi(paramsMap["throws"]); err != nil {
			fmt.Printf("Error %s", err)
		}
	} else {
		throws = 1
	}
	if dice, err = strconv.Atoi(paramsMap["dice"]); err != nil {
		fmt.Println("Error")
	}

	db := DiceBag{}
	db.result = make([]int, throws)
	d := Dice{dice}

	if len(paramsMap["exploding"]) > 0 {
		exploding = true
	}

	for x := 0; x < throws; x++ {
		ok := false
		for !ok {
			throwResult := d.Throw()
			db.result[x] += throwResult
			if exploding && throwResult == dice {
				ok = false
			} else {
				ok = true
			}
		}

	}

	if len(paramsMap["modifier"]) > 0 {
		var modValue int
		if modValue, err = strconv.Atoi(paramsMap["modValue"]); err != nil {
			fmt.Println("Error modValue")
		}
		if paramsMap["modifier"] == "-" {

			modValue *= -1
		}
		db.modifier = modValue
	}
	return db
}

// Sum returns sum of results
func (db DiceBag) Sum() int {
	result := 0
	for x := 0; x < len(db.result); x++ {
		result += db.result[x]
	}
	result += db.modifier
	return result
}

// GetResult returns throw as array
func (db DiceBag) GetResult() []int {
	return db.result
}

// Throw simple dice throw
func (d Dice) Throw() int {
	initialize()
	return rand.Intn(d.Sides) + 1
}
