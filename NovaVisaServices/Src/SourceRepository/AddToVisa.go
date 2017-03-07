package SourceRepository

import(	CR "ConfigRepository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type addToDB interface {
	AddUserToDB(CR.User)string
	CreateAnnouncements(newAnn CR.Announcement) string
}
type addingToDB struct{}

func AddToDB() addToDB{
	return &addingToDB{}
}

func (a addingToDB) AddUserToDB(newUser CR.User) string{
	var statusMsg string
	var existingUser CR.User
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		statusMsg = "Error connecting to DB"
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	userColl := session.DB(CR.DBInstance).C(CR.UserMasterColl)

	err = userColl.Find(bson.M{"email":newUser.Email}).One((&existingUser))
	if(len(existingUser.Email)==0){
		fr := FuncRep()
		validData := fr.validateData(newUser)
		if(validData == "valid") {
			pswd := fr.createPswd(5)
			err = userColl.Insert(&CR.UserCollStruct{Fname:newUser.Fname, MI:newUser.MI, Lname:newUser.Lname, Email:newUser.Email,
				Pswd:pswd, IsActive:newUser.IsActive})
			if err != nil {
				statusMsg = "Error inserting user in DB"
				panic(err)
			} else {
				statusMsg = "User added to DB"
			}
		}else {return validData}
	}else {return "user exists"}
	return statusMsg
}

func (a addingToDB) CreateAnnouncements(newAnn CR.Announcement) string{
	var statusMsg string
	var maxId CR.Announcement
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		statusMsg = "Error connecting to DB"
		panic(err)
	}
	defer session.Close()
	annColl := session.DB(CR.DBInstance).C(CR.AnnouncementsColl)

	annColl.Find(bson.M{}).Sort("-annid").Limit(1).One(&maxId)
	newAnn.AnnID = maxId.AnnID+1
	err = annColl.Insert(&CR.Announcement{AnnID:newAnn.AnnID, AnnTitle:newAnn.AnnTitle,
		AnnText:newAnn.AnnText, AnnDate:newAnn.AnnDate, AnnActive:newAnn.AnnActive})
	if err!= nil{
		statusMsg = "Error connecting to DB"
		panic(err)
	}else{ statusMsg = "created new announcement"}
	return statusMsg
}