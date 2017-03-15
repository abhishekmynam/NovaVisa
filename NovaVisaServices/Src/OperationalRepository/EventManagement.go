package OperationalRepository

import (SR "SourceRepository"
	CR "ConfigRepository"
)
type EventManagement interface{
	NewEventAddition(newEvent CR.Events, eventDesc CR.EventDetails)string
	EventUpdate (newEvent CR.Events, eventDesc CR.EventDetails)string
	EventDeactivate(newEvent CR.Events)string
	PostAPoll(newPoll CR.Polling) string
	PostVote (pollId int, itemId int)string
	PostComment(eventId int, comment string )string
	GetAllEvents()[]CR.Events
	GetActiveEvents()[]CR.Events
	GetEventDesc(eventId int)CR.EventDetails
	GetComments(eventId int)CR.EventComments
	GetPollResults(pollId int)CR.Polling
	GetPollList()[]CR.Polling
}
type eventMgmt struct{}

func EventManage() EventManagement{
	return &eventMgmt{}
}

func (e eventMgmt) NewEventAddition(newEvent CR.Events, eventDesc CR.EventDetails)string{
	var statusMsg string
	eventId := SR.AddToDB().CreateNewEvent(newEvent)
	eventDesc.EventId = eventId
	statusMsg = SR.AddToDB().CreateEventDesc(eventDesc)
	return statusMsg
}

func (e eventMgmt)EventUpdate (newEvent CR.Events, eventDesc CR.EventDetails)string{
	var statusMsg string
	statusMsg = SR.UpdateData().UpdateEvent(newEvent)
	if (eventDesc.EventId == 0 && statusMsg !="Error updating event"){
		statusMsg = SR.UpdateData().UpdateEventDesc(eventDesc)
	}
	return statusMsg
}

func (e eventMgmt)EventDeactivate(newEvent CR.Events)string{
	var statusMsg string
	newEvent.EventActive = false
	statusMsg = SR.UpdateData().UpdateEvent(newEvent)
	return statusMsg
}

func(e eventMgmt)GetAllEvents()[]CR.Events{
	var eventList []CR.Events
	eventList = SR.GetFromDB().GetAllEventsList()
	return eventList
}

func (e eventMgmt) GetActiveEvents()[]CR.Events{
	var eventList []CR.Events
	eventList = SR.GetFromDB().GetActiveEventsList()
	return eventList
}

func (e eventMgmt) GetEventDesc(eventId int)CR.EventDetails{
	var eventDesc CR.EventDetails
	eventDesc = SR.GetFromDB().GetEventDetails(eventId)
	return eventDesc
}

func (e eventMgmt) PostComment(eventId int, comment string )string{
	var status string
	status = SR.AddToDB().CreateComments(comment, eventId)
	return status
}

func (e eventMgmt) GetComments(eventId int)CR.EventComments{
	var comments CR.EventComments
	comments = SR.GetFromDB().GetEventComments(eventId)
	return comments
}

func (e eventMgmt) PostAPoll(newPoll CR.Polling) string{
	var status string
	status = SR.AddToDB().CreatePoll(newPoll)
	return status
}

func (e eventMgmt) PostVote (pollId int, itemId int)string{
	var status string
	status = SR.UpdateData().UpdatePoll(pollId, itemId)
	return status
}

func (e eventMgmt) GetPollResults(pollId int)CR.Polling{
	var pollResults CR.Polling
	pollResults = SR.GetFromDB().GetPollResults(pollId)
	return pollResults
}

func (e eventMgmt) GetPollList()[]CR.Polling{
	var pollResults []CR.Polling
	pollResults = SR.GetFromDB().GetPollList()
	return pollResults
}