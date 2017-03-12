package SourceRepository

import(	CR "ConfigRepository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type addToDB interface {
	AddUserToDB(CR.User)string
	CreateAnnouncements(newAnn CR.Announcement) string
	CreateNewEvent (newEvent CR.Events) int
	CreateEventDesc(eventDetail CR.EventDetails) string
	CreateComments(comment string, eventId int)string
	CreatePoll (poll CR.Polling) string
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
		statusMsg = "Error creating announcement"
		panic(err)
	}else{ statusMsg = "created new announcement"}
	return statusMsg
}


func (a addingToDB) CreateNewEvent (newEvent CR.Events) int{
	var maxEventId CR.Events
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	eventColl := session.DB(CR.DBInstance).C(CR.EventsColl)

	eventColl.Find(bson.M{}).Sort("-eventid").Limit(1).One(&maxEventId)
	newEvent.EventId = maxEventId.EventId+1

	err = eventColl.Insert(&CR.Events{EventId:newEvent.EventId, EventTitle:newEvent.EventTitle, EventDate:newEvent.EventDate,
				EventPostDate:newEvent.EventPostDate, EventCutOffDate:newEvent.EventCutOffDate, EventActive:newEvent.EventActive})

	if err!= nil{
		panic(err)
	}
	return maxEventId.EventId+1

}

func (a addingToDB) CreateEventDesc(eventDetail CR.EventDetails) string{
	var statusMsg string
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	eventColl := session.DB(CR.DBInstance).C(CR.EventDetailColl)
	err = eventColl.Insert(&CR.EventDetails{EventId:eventDetail.EventId,EventDetail:eventDetail.EventDetail,EventPollId:eventDetail.EventPollId,
				EventLiveStreaming:eventDetail.EventLiveStreaming, EventLiveStreamingLink:eventDetail.EventLiveStreamingLink})

	if err!= nil{
		statusMsg = "Error creating event desc"
		panic(err)
	}else{ statusMsg = "created event desc"}
	return statusMsg
}

func (a addingToDB) CreateComments(comment string, eventId int)string{
	var statusMsg string
	var comm CR.EventComments
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	eventCommColl := session.DB(CR.DBInstance).C(CR.EventCommentColl)

	err = eventCommColl.Find(bson.M{"eventid":eventId}).One(&comm)

	if(comm.EventId ==0){
		thisComment := []string{comment}
		err = eventCommColl.Insert(&CR.EventComments{EventId:eventId,EventComments:thisComment })
	}else{
		err = eventCommColl.Update(bson.M{"eventid":eventId},bson.M{"$push":bson.M{"eventcomments":comment}})
	}

	if err!= nil{
		statusMsg = "Error posting comments"
		panic(err)
	}else{ statusMsg = "posted comments"}
	return statusMsg
}

func (a addingToDB) CreatePoll (poll CR.Polling) string{
	var statusMsg string
	var polls CR.Polling
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	pollColl := session.DB(CR.DBInstance).C(CR.PollColl)
	err = pollColl.Find(bson.M{}).Sort("-pollid").Limit(1).One(&polls)
	poll.PollingId = polls.PollingId+1

	err = pollColl.Insert(&CR.Polling{PollingId:poll.PollingId, PollingName:poll.PollingName,
					PollingItems:poll.PollingItems})
	if err!= nil{
		statusMsg = "Error posting poll"
		panic(err)
	}else{ statusMsg = "posted poll"}
	return statusMsg
}
































