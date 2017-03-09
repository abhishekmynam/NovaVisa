package OperationalRepository

import (SR "SourceRepository"
	CR "ConfigRepository"
)
type EventManagement interface{
	NewEventAddition(newEvent CR.Events)string
	EventUpdate (newEvent CR.Events)string
	EventDeactivate(newEvent CR.Events)string
	GetAllEvents()[]CR.Events
	GetActiveEvents()[]CR.Events
}
type eventMgmt struct{}

func EventManage() UserManagement{
	return &userMgmt{}
}

func (e eventMgmt) NewEventAddition(newEvent CR.Events){
	var statusMsg string
	statusMsg = SR.AddToDB().CreateNewEvent(newEvent)
	return statusMsg
}

func (e eventMgmt)EventUpdate (newEvent CR.Events){
	var statusMsg string
	statusMsg = SR.UpdateData().UpdateEvent(newEvent)
	return statusMsg
}

func (e eventMgmt)EventDeactivate(newEvent CR.Events){
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