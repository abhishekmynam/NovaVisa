package ConfigRepository


type User struct {
	Fname string
	MI string
	Lname string
	Email string
	Pswd string
	ConfPswd string
	IsActive bool
}

type UserCollStruct struct {
	Fname string
	MI string
	Lname string
	Email string
	Pswd string
	IsActive bool
}

type Announcement struct{
	AnnID int
	AnnTitle string
	AnnText string
	AnnDate string
	AnnActive bool
}

type Events struct {
	EventId int
	EventTitle string
	EventPostDate string
	EventDate string
	EventCutOffDate string
	EventActive bool
}

func BoolToByte(str bool) byte {
	switch str {
		case true:
			return 1
		case false:
			return 0
	}
	return 0
}