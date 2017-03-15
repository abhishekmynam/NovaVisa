package SourceRepository

import(	CR "ConfigRepository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type getFromDB interface {
	GetUserList()[]CR.UserCollStruct
	GetAnnouncementList() []CR.Announcement
	GetAllAnnouncementList() []CR.Announcement
	GetAuthPswd (email string)string
	GetActiveEventsList() []CR.Events
	GetAllEventsList() []CR.Events
	GetEventDetails(eventId int) CR.EventDetails
	GetEventComments(eventId int) CR.EventComments
	GetPollResults(pollId int) CR.Polling
	GetPollList()[]CR.Polling
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

func (g gettingFromDB) GetAllEventsList() []CR.Events{
	var eventList []CR.Events
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	eventColl := session.DB(CR.DBInstance).C(CR.EventsColl)

	err = eventColl.Find(bson.M{}).Select(nil).Sort("eventid").All(&eventList)
	if err!= nil{
		panic(err)
	}
	return eventList

}

func (g gettingFromDB) GetActiveEventsList() []CR.Events{
	var eventList []CR.Events
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	eventColl := session.DB(CR.DBInstance).C(CR.EventsColl)

	err = eventColl.Find(bson.M{"eventactive":true,"eventdate":bson.M{"$gte":time.Now().Format("20060102150405")}}).Select(nil).Sort("eventid").All(&eventList)
	if err!= nil{
		panic(err)
	}
	return eventList
}

func (g gettingFromDB) GetEventDetails(eventId int) CR.EventDetails{
	var eventDetail CR.EventDetails
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	eventColl := session.DB(CR.DBInstance).C(CR.EventDetailColl)
	err = eventColl.Find(bson.M{"eventid":eventId}).Select(nil).One(&eventDetail)
	if err!= nil{
		panic(err)
	}
	return  eventDetail
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

func (g gettingFromDB) GetEventComments(eventId int) CR.EventComments{
	var comments CR.EventComments
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	eventColl := session.DB(CR.DBInstance).C(CR.EventCommentColl)

	err = eventColl.Find(bson.M{"eventid":eventId}).One(&comments)
	if err!= nil{
		panic(err)
	}
	return comments
}

func (g gettingFromDB) GetPollResults(pollId int) CR.Polling{
	var polls CR.Polling
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	pollColl := session.DB(CR.DBInstance).C(CR.PollColl)

	err = pollColl.Find(bson.M{"pollingid":pollId}).One(&polls)
	if err!= nil{
		panic(err)
	}
	return polls
}

func (g gettingFromDB)GetPollList()[]CR.Polling{
	var polls []CR.Polling
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	pollColl := session.DB(CR.DBInstance).C(CR.PollColl)
	err = pollColl.Find(bson.M{}).Select(bson.M{"pollingid":1,"pollingname":1}).Sort("-pollingid").All(&polls)
	if err!= nil{
		panic(err)
	}
	return polls
}