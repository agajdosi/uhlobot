package generate

import (
	"math/rand"
	"strconv"
)

//Password generates random password
func Password(n int) string {
	var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//Username generates random username
func Username(n int) string {
	var letters = []rune("1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	names := []string{
		"jiri",
		"jan",
		"honza",
		"petr",
		"daniel",
		"andrej",
		"adam",
		"filip",
		"pavel",
	}

	username := names[rand.Intn(len(names))] + string(b)

	return username
}

//Nickname creates a funky alias!
func Nickname() string {
	aliases := []string{
		"Pohoda",
		"Tramp",
		"Kyno",
		"Subaru",
		"BMW",
		"Bastl",
		"Zlato",
		"Love",
		"Silver",
	}

	alias := aliases[rand.Intn(len(aliases))] + strconv.Itoa(rand.Intn(100))
	return alias
}
