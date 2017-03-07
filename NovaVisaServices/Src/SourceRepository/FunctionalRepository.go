package SourceRepository


import(	"crypto/rand"
	"fmt"
	"ConfigRepository"
	"regexp"
)

type funcRep interface {
	createPswd(n int)string
	validateData(ConfigRepository.User)string
	EncryptPswd (pswd string)string
}
type functionalRep struct{}

func FuncRep() funcRep{
	return &functionalRep{}
}

func (f functionalRep)createPswd(n int)string{
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

func (f functionalRep)EncryptPswd (pswd string)string{
	var newPswd string
	return newPswd
}

func (f functionalRep) validateData(user ConfigRepository.User)string{
	dataValid := "valid"
	isAlpha := regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString
	if(user.Fname!= ""){
		if !isAlpha(user.Fname) {
			return "invalid first name"
		}
	}else if(user.Fname== ""){return "missing first name"}
	if(user.Lname!= ""){
		if !isAlpha(user.Lname) {
			return "invalid last name"
		}
	}else if(user.Lname== ""){return "missing last name"}
	if(user.MI!= ""){
		if !isAlpha(user.MI) {
			return "invalid middle initial"
		}
	}else if(user.MI== ""){return "missing middle initial"}
	return dataValid
}