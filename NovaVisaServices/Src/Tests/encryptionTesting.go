package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {
	password := []byte("abhishek")
	thisPswd := []byte("$2a$10$3D6Mq6C/iHplYQ36dLf6VuPXaejW4jLxpCiAoBQPJHbWyimPK/D3.")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(thisPswd, password)
	fmt.Println(err) // nil means it is a match
}