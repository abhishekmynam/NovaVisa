package main

import ("fmt"
	CR "ConfigRepository"
	SR "SourceRepository"
	"regexp"
	"crypto/rand"
	"time"
	"github.com/google/flatbuffers/go"
	fb "FlatBuffers/FlatBufs"
	OR "OperationalRepository"
	//"bytes"
	//"encoding/binary"
)

func main() {
	//fmt.Println(CR.DBserver)
	//pswdGen()
	//validateData()
	//createUser()
	//updateUser()
	//createAnnouncement()
	//updateAnnouncement()
	//getUsers()
	//getannouncements()
	//createEvent()
	//updateEvent()
	//getActiveEvents()
	//createEventDesc()
	//getEventdesc()
	//createComment()
	//getComment()
	//polls()
	//testingThis()
	pollsRelated()
}
func testingThis()CR.Polling{
	var poll CR.Polling
	var pollEntry = []CR.PollingEntries{
		CR.PollingEntries{
			ItemId: 1,
			Item: "Item 1",
			Votes :0,
		},
		CR.PollingEntries{
			ItemId: 2,
			Item: "Item 2",
			Votes :0,
		},
	}

	poll.PollingItems = pollEntry
	poll.PollingName = "first poll"
	fmt.Println(poll)
	return poll

}

func pollsRelated(){
	var Thispoll CR.Polling
	var ThispollEntry = []CR.PollingEntries{
		CR.PollingEntries{
			ItemId: 1,
			Item: "Item 1",
			Votes :0,
		},
		CR.PollingEntries{
			ItemId: 2,
			Item: "Item 2",
			Votes :0,
		},
	}

	Thispoll.PollingItems = ThispollEntry
	Thispoll.PollingName = "first poll"
	//statusMsg := OR.EventManage().PostAPoll(Thispoll)
	fmt.Println(Thispoll)
	statusMsg := OR.EventManage().PostVote(1,1)
	thisObject := OR.EventManage().GetPollResults(1)
	fmt.Println(statusMsg)
	fmt.Println(thisObject)
	fmt.Println(statusMsg)
}


func getComment(){
	comment := SR.GetFromDB().GetEventComments(2)
	fmt.Println(comment)
}

func createComment(){
	event := SR.AddToDB().CreateComments("this comment 1",2)
	fmt.Println(event)
}

func getEventdesc(){
	event := SR.GetFromDB().GetEventDetails(3)
	fmt.Println(event)
}

func createEventDesc(){
	var event CR.Events
	event.EventTitle = "second event"
	event.EventPostDate =time.Now().Format("20060102150405")
	event.EventDate = "20170409100000"
	event.EventCutOffDate="20170409090000"
	event.EventActive = true
	var eventDetaill CR.EventDetails
	eventDetaill.EventPollId = 1
	eventDetaill.EventDetail = "this is an event"
	//b := []string{"Penn", "Teller"}
	//eventDetaill.EventComments =b
	eventDetaill.EventLiveStreaming = true
	eventDetaill.EventLiveStreamingLink = "www.google.com"
	thismsg := OR.EventManage().NewEventAddition(event,eventDetaill)
	fmt.Println(thismsg)
}
func getActiveEvents(){
	lists := SR.GetFromDB().GetActiveEventsList();
	fmt.Println(lists)
}

func createEvent(){
	var event CR.Events
	event.EventTitle = "first event"
	event.EventPostDate =time.Now().Format("20060102150405")
	event.EventDate = "20170409100000"
	event.EventCutOffDate="20170409090000"
	event.EventActive = true
	msg := SR.AddToDB().CreateNewEvent(event)
	fmt.Println(msg)
}

func updateEvent(){
	var event CR.Events
	event.EventTitle = "first event"
	event.EventPostDate =time.Now().Format("20060102150405")
	event.EventDate = "20170409100000"
	event.EventCutOffDate="20170409090000"
	event.EventActive = true
	event.EventId = 1
	msg := SR.UpdateData().UpdateEvent(event)
	fmt.Println(msg)
}
func getUsers(){
	userList := SR.GetFromDB().GetUserList()
	fmt.Println(userList)
}

func getannouncements(){
	annList := SR.GetFromDB().GetAllAnnouncementList()
	fmt.Println(annList)
	builder := flatbuffers.NewBuilder(1024)
	var thisobjlist [12] flatbuffers.UOffsetT
	for i,j := range annList{
		annTitle := builder.CreateString(j.AnnTitle)
		annText := builder.CreateString(j.AnnText)
		annDate := builder.CreateString(j.AnnDate)
		fb.AnnouncementStart(builder)
		fb.AnnouncementAddAnnId(builder,int32(j.AnnID))
		fb.AnnouncementAddAnnTitle(builder,annTitle)
		fb.AnnouncementAddAnnText(builder, annText)
		fb.AnnouncementAddAnnActive(builder,CR.BoolToByte(j.AnnActive))
		fb.AnnouncementAddAnnDate(builder,annDate)
		thisobj:= fb.AnnouncementEnd(builder)
		thisobjlist[i] = thisobj
	}
	fb.AnnouncementListStartAnnListVector(builder,len(annList))
	for _,j:=range thisobjlist{
		builder.PlaceUOffsetT(j)
	}
	finalObj := fb.AnnouncementListEnd(builder)
	//annNames :=  builder.CreateString("thisname")
	//fb.AnnouncementListAddAnnName(builder,annNames)
	builder.Finish(finalObj)
	buf:= builder.FinishedBytes()
	fmt.Println(buf)
	/*bufItem := new(bytes.Buffer)
	binary.Write(bufItem, binary.LittleEndian, buf)
	buf1 := bufItem.Bytes()
	buffyRead := bytes.NewReader(buf1)
	var buffy []byte
	binary.Read(buffyRead, binary.LittleEndian, &buffy)*/
	monster :=  fb.GetRootAsAnnouncementList(buf,0)
	anns := new(fb.Announcement)
	if monster.AnnList(anns,1){
		thisLists := anns.AnnTitle()
		fmt.Println(thisLists)
	}
	fmt.Println(monster)

}
func pswdGen(){
	n:=5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	fmt.Println(s)
}

func createUser(){
	var user CR.User
	user.Fname = "Abhishek123"
	user.Lname = "Mynam123"
	user.MI = "H123"
	user.Email = "aabhishek.mynam@gmail.com"
	msg := SR.AddToDB().AddUserToDB(user)
	fmt.Println(msg)
}

func validateData(){
	user := "this123"
	isAlpha := regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString
	if(user!= ""){
		if !isAlpha(user) {
			fmt.Println("invalid data")
		}
	}
}

func updateUser(){
	var user CR.User
	user.Fname = "Abhishek123"
	user.Lname = "Mynam123"
	user.MI = "HW123"
	user.Email = "abhishek.mynam123@gmail.com"
	user.Pswd = "thisPswd123"
	msg := SR.UpdateData().UpdateUser(user)
	fmt.Println(msg)
}

func createAnnouncement(){
	var ann CR.Announcement
	ann.AnnTitle ="test announcement4"
	ann.AnnText = "jumping japang 123"
	ann.AnnActive = true
	ann.AnnDate = time.Now().Format("20060102150405")
	msg := SR.AddToDB().CreateAnnouncements(ann)
	fmt.Println(msg)
}

func updateAnnouncement(){
	var ann CR.Announcement
	ann.AnnID = 2
	ann.AnnTitle ="test announcement1"
	ann.AnnText = "jumping japang IPL 123"
	ann.AnnActive = true
	ann.AnnDate = time.Now().Format("20060102150405")
	msg := SR.UpdateData().UpdateAnnouncement(ann)
	fmt.Println(msg)
}
