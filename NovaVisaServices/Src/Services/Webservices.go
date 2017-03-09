package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	//flatbuffers "github.com/google/flatbuffers/go"
	OR "OperationalRepository"
	//fb "FlatBuffers/FlatBufs"
//	CR "ConfigRepository"
)


func main() {
	http.HandleFunc("/activeannouncements",getActiveAnnouncements)
	http.ListenAndServe("localhost:1337", nil)

}

func getActiveAnnouncements(w http.ResponseWriter, r *http.Request){
	annList := OR.AnnouncementManage().GetAllAnnouncements()
	/*builder := flatbuffers.NewBuilder(1024)

	for _,j := range annList{
		//AnnName := builder.CreateString(j.AnnTitle)
		fb.AnnouncementStart(builder)
		fb.AnnouncementAddAnnId(builder,int32(j.AnnID))
		fb.AnnouncementAddAnnTitle(builder,builder.CreateString(j.AnnTitle))
		fb.AnnouncementAddAnnText(builder, builder.CreateString(j.AnnText))
		fb.AnnouncementAddAnnActive(builder,CR.BoolToByte(j.AnnActive))
		fb.AnnouncementAddAnnDate(builder,builder.CreateString(j.AnnDate))
		builder.PrependUOffsetT(builder.CreateString(j.AnnTitle))
	}
	AnnouncementList := builder.EndVector(len(annList))
	fmt.Println(string(AnnouncementList))*/
	jsonAnnList,_ := json.Marshal(annList)
	fmt.Fprintf(w,string(jsonAnnList))
}


