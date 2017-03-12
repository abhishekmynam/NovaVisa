package SourceRepository

import (
	CR "ConfigRepository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type updateDB interface {
	UpdateUser(newUser CR.User)string
	UpdateAnnouncement (newAnn CR.Announcement)string
	UpdateEvent (newEvent CR.Events)string
	UpdateEventDesc(eventDesc CR.EventDetails)string
	UpdatePoll(pollId int, itemId int)string
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

func (u updatingDB)UpdateEvent (newEvent CR.Events)string{
	var statusMsg string
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		statusMsg = "Error connecting to DB"
		panic(err)
	}
	defer session.Close()
	eventColl := session.DB(CR.DBInstance).C(CR.EventsColl)

	err = eventColl.Update(bson.M{"eventid":newEvent.EventId},&CR.Events{EventId:newEvent.EventId, EventTitle:newEvent.EventTitle, EventDate:newEvent.EventDate,
		EventPostDate:newEvent.EventPostDate, EventCutOffDate:newEvent.EventCutOffDate, EventActive:newEvent.EventActive})

	if err!= nil{
		statusMsg = "Error updating event"
		panic(err)
	}else{ statusMsg = "updated event"}
	return statusMsg
}

func (u updatingDB)UpdateEventDesc(eventDesc CR.EventDetails)string{
	var statusMsg string
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		statusMsg = "Error connecting to DB"
		panic(err)
	}
	defer session.Close()
	eventColl := session.DB(CR.DBInstance).C(CR.EventDetailColl)
	err = eventColl.Update(bson.M{"eventid":eventDesc.EventId},&CR.EventDetails{EventId:eventDesc.EventId, EventLiveStreaming:eventDesc.EventLiveStreaming,
		EventLiveStreamingLink:eventDesc.EventLiveStreamingLink, EventDetail:eventDesc.EventDetail,
		EventPollId:eventDesc.EventPollId})
	if err!= nil{
		statusMsg = "Error updating event"
		panic(err)
	}else{ statusMsg = "updated event"}
	return statusMsg
}

func (u updatingDB)UpdatePoll(pollId int, itemId int)string{
	var statusMsg string
	var thisPoll CR.Polling
	session, err:= mgo.Dial(CR.DBserver)
	if err!= nil{
		statusMsg = "Error connecting to DB"
		panic(err)
	}
	defer session.Close()
	pollColl := session.DB(CR.DBInstance).C(CR.PollColl)

	err = pollColl.Find(bson.M{"pollid":pollId}).One(&thisPoll)
	for i,j := range thisPoll.PollingItems{
		if(j.ItemId == itemId){
			thisPoll.PollingItems[i].Votes = thisPoll.PollingItems[i].Votes+1
		}
	}
	err = pollColl.Update(bson.M{"pollid":pollId},&CR.Polling{PollingId:pollId, PollingName:thisPoll.PollingName,
					PollingItems:thisPoll.PollingItems})
	if err!= nil{
		statusMsg = "Error polling"
		panic(err)
	}else{ statusMsg = "voted"}
	return statusMsg
}