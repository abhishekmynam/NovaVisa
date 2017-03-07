package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	OR "OperationalRepository"
)


func main() {
	http.HandleFunc("/activeannouncements",getActiveAnnouncements)
	http.ListenAndServe("localhost:1337", nil)

}

func getActiveAnnouncements(w http.ResponseWriter, r *http.Request){
	annList := OR.AnnouncementManage().GetAllAnnouncements()
	jsonAnnList,_ := json.Marshal(annList)
	fmt.Fprintf(w,string(jsonAnnList))
}
