package main

import (
	OR "OperationalRepository"
	CR "ConfigRepository"
	"fmt"
	"time"
)

func main(){
	//addUser()
	//updateUsers()
	//deactivateUsers()
	//allUsers()
	//createAnn()
	//updateAnn()
	//getActiveAnn()
	//getAllAnn()
	authUser()
	}

func getActiveAnn(){
	anns:= OR.AnnouncementManage().GetActiveAnnouncements()
	fmt.Println(anns)
}

func authUser(){
	user := OR.UserManage().AuthUser("alimynam@gmail.com","password")
	fmt.Println(user)
}

func getAllAnn(){
	anns:= OR.AnnouncementManage().GetAllAnnouncements()
	fmt.Println(anns)
}

func updateAnn (){
	var ann CR.Announcement
	ann.AnnID = 2
	ann.AnnTitle ="test announcement1"
	ann.AnnText = "jumping japang 123"
	ann.AnnActive = false
	ann.AnnDate = time.Now().Format("20060102150405")
	ann.AnnActive = true
	msg := OR.AnnouncementManage().AnnouncementUpdate(ann)
	fmt.Println(msg)
}

func createAnn (){
	var ann CR.Announcement
	ann.AnnTitle ="test announcement1"
	ann.AnnText = "jumping japang 123"
	ann.AnnActive = false
	ann.AnnDate = time.Now().Format("20060102150405")
	ann.AnnActive = true
	msg := OR.AnnouncementManage().NewAnnouncementAdd(ann)
	fmt.Println(msg)
}
func addUser(){
	var user CR.User
	user.Fname = "Ali"
	user.Lname = "Mynam"
	user.MI = "H"
	user.Email = "alimynam1@gmail.com"
	user.IsActive = true
	msg := OR.UserManage().NewUserAddition(user)
	fmt.Println(msg)
}


func updateUsers(){
	var user CR.User
	user.Fname = "Ali"
	user.Lname = "Mynam"
	user.MI = "H"
	user.Pswd = "password"
	user.Email = "alimynam@gmail.com"
	user.IsActive = true
	msg := OR.UserManage().UserUpdate(user)
	fmt.Println(msg)
}

func deactivateUsers(){
	var user CR.User
	user.Fname = "Ali"
	user.Lname = "Mynam"
	user.MI = "H"
	user.Pswd = "password"
	user.Email = "alimynam@gmail.com"
	user.IsActive = true
	msg := OR.UserManage().UserDeactivate(user)
	fmt.Println(msg)
}

func allUsers(){
	msg := OR.UserManage().GetAllUsers()
	fmt.Println(msg)
}