package main

import ("fmt"
	CR "ConfigRepository"
	SR "SourceRepository"
	"regexp"
	"crypto/rand"
	"time"
)

func main() {
	fmt.Println(CR.DBServerTest)
	//pswdGen()
	//validateData()
	//createUser()
	//updateUser()
	//createAnnouncement()
	//updateAnnouncement()
	//getUsers()
	getannouncements()
}

func getUsers(){
	userList := SR.GetFromDB().GetUserList()
	fmt.Println(userList)
}

func getannouncements(){
	annList := SR.GetFromDB().GetAnnouncementList()
	fmt.Println(annList)
}
func pswdGen(){
	n:=5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	fmt.Println(s)
}

func createUser(){
	var user CR.User
	user.Fname = "Abhishek123"
	user.Lname = "Mynam123"
	user.MI = "H123"
	user.Email = "abhishek.mynam123@gmail.com"
	msg := SR.AddToDB().AddUserToDB(user)
	fmt.Println(msg)
}

func validateData(){
	user := "this123"
	isAlpha := regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString
	if(user!= ""){
		if !isAlpha(user) {
			fmt.Println("invalid data")
		}
	}
}

func updateUser(){
	var user CR.User
	user.Fname = "Abhishek123"
	user.Lname = "Mynam123"
	user.MI = "HW123"
	user.Email = "abhishek.mynam123@gmail.com"
	user.Pswd = "thisPswd123"
	msg := SR.UpdateData().UpdateUser(user)
	fmt.Println(msg)
}

func createAnnouncement(){
	var ann CR.Announcement
	ann.AnnTitle ="test announcement1"
	ann.AnnText = "jumping japang 123"
	ann.AnnActive = true
	ann.AnnDate = time.Now().Format("20060102150405")
	msg := SR.AddToDB().CreateAnnouncements(ann)
	fmt.Println(msg)
}

func updateAnnouncement(){
	var ann CR.Announcement
	ann.AnnID = 2
	ann.AnnTitle ="test announcement1"
	ann.AnnText = "jumping japang IPL 123"
	ann.AnnActive = true
	ann.AnnDate = time.Now().Format("20060102150405")
	msg := SR.UpdateData().UpdateAnnouncement(ann)
	fmt.Println(msg)
}
