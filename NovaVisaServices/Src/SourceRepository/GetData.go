package SourceRepository

import(	CR "ConfigRepository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type getFromDB interface {
	GetUserList()[]CR.UserCollStruct
	GetAnnouncementList() []CR.Announcement
	GetAllAnnouncementList() []CR.Announcement
	GetAuthPswd (email string)string
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

	err = userColl.Find(nil).Select(bson.M{"fname":1,"mi":1,"lname":1,"email":1,"isactive":1}).Sort("email").All(&userList)
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

	err = annColl.Find(bson.M{"annactive":true}).Select(nil).Sort("annid").All(&anncmtList)
	if err!= nil{
		panic(err)
	}
	return anncmtList
}

func (g gettingFromDB) GetAllAnnouncementList() []CR.Announcement{
	var anncmtList []CR.Announcement

	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	annColl := session.DB(CR.DBInstance).C(CR.AnnouncementsColl)

	err = annColl.Find(bson.M{}).Select(nil).Sort("annid").All(&anncmtList)
	if err!= nil{
		panic(err)
	}
	return anncmtList
}

func (g gettingFromDB) GetAuthPswd (email string)string{
	var pswd CR.UserCollStruct

	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	userColl := session.DB(CR.DBInstance).C(CR.UserMasterColl)

	err = userColl.Find(bson.M{"email":email}).One(&pswd)
	if err!= nil{
		panic(err)
	}
	return pswd.Pswd
}