package main

import (
	"log"

	"github.com/akselarzuman/go-utils"
)

type User struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	users := []User{
		{
			Name:    "Aksel",
			Surname: "Arzuman",
			Age:     30,
		},
		{
			Name:    "Zeynep",
			Surname: "Arzuman",
			Age:     25,
		},
		{
			Name:    "John",
			Surname: "Smith",
			Age:     25,
		},
	}

	// Find the users with the surname "Arzuman"
	usersFiltered := utils.Filter(users, func(u User) bool {
		return u.Surname == "Arzuman"
	})

	log.Println(usersFiltered)

	// Find the users with the surname "Arzuman" and the age 25
	usersFiltered = utils.Filter(users, func(u User) bool {
		return u.Surname == "Arzuman" && u.Age == 25
	})

	log.Println(usersFiltered)

	// Check if there is any user with the surname "Smith"
	if ok := utils.Any(users, func(u User) bool {
		return u.Surname == "Smith"
	}); ok {
		log.Println("There is a user with the surname 'Smith'")
	} else {
		log.Println("There is no user with the surname 'Smith'")
	}
}
