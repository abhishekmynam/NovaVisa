package SourceRepository

import (
	CR "ConfigRepository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type updateDB interface {
	UpdateUser(newUser CR.User)string
	UpdateAnnouncement (newAnn CR.Announcement)string
}
type updatingDB struct{}

func UpdateData() updateDB{
	return &updatingDB{}
}

func (u updatingDB)UpdateUser(newUser CR.User)string{
	var statusMsg string

	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		statusMsg = "Error connecting to DB"
		panic(err)
	}
	defer session.Close()
	userColl := session.DB(CR.DBInstance).C(CR.UserMasterColl)
	fr := FuncRep()
	validData := fr.validateData(newUser)
	if(validData == "valid") {
		err = userColl.Update(bson.M{"email":newUser.Email},&CR.UserCollStruct{Fname:newUser.Fname, MI:newUser.MI, Lname:newUser.Lname,
			Email:newUser.Email, Pswd:newUser.Pswd, IsActive:newUser.IsActive})
		if err != nil {
			statusMsg = "Error updating user in DB"
			panic(err)
		} else {
			statusMsg = "User updated in DB"
		}
	}else {return validData}

	return statusMsg
}

func (u updatingDB)UpdateAnnouncement (newAnn CR.Announcement)string{
	var statusMsg string
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		statusMsg = "Error connecting to DB"
		panic(err)
	}
	defer session.Close()
	annColl := session.DB(CR.DBInstance).C(CR.AnnouncementsColl)

	err = annColl.Update(bson.M{"annid":newAnn.AnnID},&CR.Announcement{AnnID:newAnn.AnnID,AnnTitle:newAnn.AnnTitle, AnnText:newAnn.AnnText,
	AnnActive:newAnn.AnnActive,AnnDate:newAnn.AnnDate})
	if err!= nil{
		statusMsg = "Error updating DB"
		panic(err)
	}else{statusMsg= "Announcement updated in DB"}
	return statusMsg
}