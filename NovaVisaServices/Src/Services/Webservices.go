package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	OR "OperationalRepository"
	CR "ConfigRepository"
	"strconv"
)


func main() {
	http.HandleFunc("/activeannouncements",getActiveAnnouncements)
	http.HandleFunc("/allannouncements",getAllAnnouncements)
	http.HandleFunc("/allusers",getAllUsers)
	http.HandleFunc("/allevents",getAllEvents)
	http.HandleFunc("/activeevents",getActiveEvents)
	http.HandleFunc("/allpolls", getAllPolls)
	http.HandleFunc("/geteventdesc", getEventDesc)
	http.HandleFunc("/getcomments", getComments)
	http.HandleFunc("/getpollresults",getPollResults)
	http.HandleFunc("/addnewuser",addNewUser)
	http.HandleFunc("/createannouncement",createAnnouncements)
	http.HandleFunc("/createcomment",createComments)
	http.HandleFunc("/addpoll",addPoll)
	http.HandleFunc("/createevent",createEvent)
	http.HandleFunc("/updateuser",updateUser)
	http.HandleFunc("/updateannouncement",updateAnnouncement)
	http.HandleFunc("/updateevent",updateEvent)
	//http.HandleFunc("/testproj", gettestproj)
	http.ListenAndServe("localhost:8080", nil)
}
/*func gettestproj(w http.ResponseWriter, r *http.Request){
	var user CR.User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
}*/
func getActiveAnnouncements(w http.ResponseWriter, r *http.Request){
	annList := OR.AnnouncementManage().GetActiveAnnouncements()
	jsonAnnList,_ := json.Marshal(annList)
	fmt.Fprintf(w,string(jsonAnnList))
}

func getAllAnnouncements(w http.ResponseWriter, r *http.Request){
	annList := OR.AnnouncementManage().GetAllAnnouncements()
	jsonAnnList,_ := json.Marshal(annList)
	fmt.Fprintf(w,string(jsonAnnList))
}

func getAllUsers(w http.ResponseWriter, r *http.Request){
	userList := OR.UserManage().GetAllUsers()
	jsonUserList,_ := json.Marshal(userList)
	fmt.Fprintf(w,string(jsonUserList))
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	eventList := OR.EventManage().GetAllEvents()
	jsonEventList,_ := json.Marshal(eventList)
	fmt.Fprintf(w,string(jsonEventList))
}

func getActiveEvents(w http.ResponseWriter, r *http.Request) {
	eventList := OR.EventManage().GetActiveEvents()
	jsonEventList,_ := json.Marshal(eventList)
	fmt.Fprintf(w,string(jsonEventList))

}

func getAllPolls(w http.ResponseWriter, r *http.Request){
	pollList := OR.EventManage().GetPollList()
	jsonPollList,_ := json.Marshal(pollList)
	fmt.Fprintf(w,string(jsonPollList))
}

func getEventDesc (w http.ResponseWriter, r *http.Request) {
	var eventId int
	r.ParseForm()
	for key, j := range r.Form {
		if (key == "eventid") {
			eventId,_= strconv.Atoi(j[0])
		}
	}
	eventDesc := OR.EventManage().GetEventDesc(eventId)
	jsonEventDesc,_ := json.Marshal(eventDesc)
	fmt.Fprintf(w,string(jsonEventDesc))

}

func getComments (w http.ResponseWriter, r *http.Request) {
	var eventId int
	r.ParseForm()
	for key, j := range r.Form {
		if (key == "eventid") {
			eventId,_= strconv.Atoi(j[0])
		}
	}
	eventDesc := OR.EventManage().GetComments(eventId)
	jsonComments,_ := json.Marshal(eventDesc)
	fmt.Fprintf(w,string(jsonComments))

}

func getPollResults (w http.ResponseWriter, r *http.Request) {
	var pollId int
	r.ParseForm()
	for key, j := range r.Form{
		if (key == "pollid") {
			pollId,_= strconv.Atoi(j[0])
		}
	}
	eventDesc := OR.EventManage().GetPollResults(pollId)
	jsonPollRes,_ := json.Marshal(eventDesc)
	fmt.Fprintf(w,string(jsonPollRes))

}

func addNewUser (w http.ResponseWriter, r *http.Request){
	var statusMsg string
	var user CR.User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	statusMsg = OR.UserManage().NewUserAddition(user)
	fmt.Fprintf(w,string(statusMsg))

}

func createAnnouncements(w http.ResponseWriter, r *http.Request) {
	var statusMsg string
	var announcement CR.Announcement
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&announcement)
	statusMsg = OR.AnnouncementManage().NewAnnouncementAdd(announcement)
	fmt.Fprintf(w,string(statusMsg))
}

func createComments(w http.ResponseWriter, r *http.Request) {
	var eventId int
	var comment, statusMsg string
	r.ParseForm()
	for key, j := range r.Form{
		if (key == "eventid") {
			eventId,_= strconv.Atoi(j[0])
		}
		if (key == "comment") {
			comment,_= strconv.Atoi(j[0])
		}
	}
	statusMsg = OR.EventManage().PostComment(eventId,comment)
	fmt.Fprintf(w,string(statusMsg))
}

func addPoll (w http.ResponseWriter, r *http.Request){
	var statusMsg string
	var poll CR.Polling
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&poll)
	statusMsg = OR.EventManage().PostAPoll(poll)
	fmt.Fprintf(w,string(statusMsg))
}

func createEvent (w http.ResponseWriter, r *http.Request){
	var statusMsg string
	var event CR.FullEvent
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&event)
	statusMsg = OR.EventManage().NewEventAddition(event)
	fmt.Fprintf(w,string(statusMsg))
}

func updateUser (w http.ResponseWriter, r *http.Request){
	var statusMsg string
	var user CR.User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	statusMsg = OR.UserManage().UserUpdate(user)
	fmt.Fprintf(w,string(statusMsg))
}

func updateAnnouncement (w http.ResponseWriter, r *http.Request){
	var statusMsg string
	var announcement CR.Announcement
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&announcement)
	statusMsg = OR.AnnouncementManage().AnnouncementUpdate(announcement)
	fmt.Fprintf(w,string(statusMsg))
}

func updateEvent (w http.ResponseWriter, r *http.Request){
	var statusMsg string
	var event CR.FullEvent
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	json.NewDecoder(r.Body).Decode(&event)
	statusMsg = OR.EventManage().EventUpdate(event)
	fmt.Fprintf(w,string(statusMsg))
}

func updatePoll (w http.ResponseWriter, r *http.Request){
	var pollId int
	var itemId, statusMsg string
	r.ParseForm()
	for key, j := range r.Form{
		if (key == "pollid") {
			pollId,_= strconv.Atoi(j[0])
		}
		if (key == "itemid") {
			itemId,_= strconv.Atoi(j[0])
		}
	}
	statusMsg = OR.EventManage().PostVote(pollId,itemId)
	fmt.Fprintf(w,string(statusMsg))
}

/*
	**UpdateUser(newUser CR.User)string
	**UpdateAnnouncement (newAnn CR.Announcement)string
	**UpdateEvent (newEvent CR.Events)string
	**UpdateEventDesc(eventDesc CR.EventDetails)string
	UpdatePoll(pollId int, itemId int)string
*/


