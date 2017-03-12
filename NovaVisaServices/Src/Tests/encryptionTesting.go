package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {

	password := []byte("abhishek")
	passwordInService, err := bcrypt.GenerateFromPassword(password, 5)
	fmt.Println(passwordInService)
	pswdInDB := DBPassword()
	err = bcrypt.CompareHashAndPassword(pswdInDB, passwordInService)
	fmt.Println(err)

}


func DBPassword()[]byte{
	password := []byte("abhishek")
	passwordInService, _ := bcrypt.GenerateFromPassword(password, 5)
	fmt.Println(passwordInService)
	pswdInDB,_ := bcrypt.GenerateFromPassword(passwordInService,5)
	fmt.Println(pswdInDB)
	return pswdInDB
}

/*	password := []byte("abhishek")

	passwordInService, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	pswdInDB,err := bcrypt.GenerateFromPassword(passwordInService, bcrypt.DefaultCost)


	if err != nil {
		panic(err)
	}
	fmt.Println(string(passwordInService))
	fmt.Println(string(pswdInDB))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(pswdInDB, passwordInService)
	fmt.Println(err) // nil means it is a match
}*/