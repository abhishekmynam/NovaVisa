package SourceRepository

import(	CR "ConfigRepository"
	"gopkg.in/mgo.v2"
)

type getFromDB interface {
	GetUserList()[]CR.UserCollStruct
	GetAnnouncementList() []CR.Announcement
}
type gettingFromDB struct{}

func GetFromDB() getFromDB{
	return &gettingFromDB{}
}

func (g gettingFromDB) GetUserList()[]CR.UserCollStruct{

	var userList []CR.UserCollStruct
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	userColl := session.DB(CR.DBInstance).C(CR.UserMasterColl)

	err = userColl.Find(nil).All(&userList)
	if err!= nil{
		panic(err)
	}
	return userList
}


func (g gettingFromDB) GetAnnouncementList() []CR.Announcement{
	var anncmtList []CR.Announcement

	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	annColl := session.DB(CR.DBInstance).C(CR.AnnouncementsColl)

	err = annColl.Find(nil).All(&anncmtList)
	return anncmtList
}