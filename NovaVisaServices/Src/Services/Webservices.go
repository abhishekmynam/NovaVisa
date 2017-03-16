package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	//flatbuffers "github.com/google/flatbuffers/go"
	OR "OperationalRepository"
	//fb "FlatBuffers/FlatBufs"
	//CR "ConfigRepository"
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
	http.ListenAndServe("localhost:8080", nil)
}

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
	for key, j := range r.Form {
		if (key == "pollid") {
			pollId,_= strconv.Atoi(j[0])
		}
	}
	eventDesc := OR.EventManage().GetPollResults(pollId)
	jsonPollRes,_ := json.Marshal(eventDesc)
	fmt.Fprintf(w,string(jsonPollRes))

}
