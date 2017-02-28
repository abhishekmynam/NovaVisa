package ConfigRepository


type User struct {
	Fname string
	MI string
	Lname string
	Email string
	Pswd string
	ConfPswd string
}

type UserCollStruct struct {
	Fname string
	MI string
	Lname string
	Email string
	Pswd string
}

type Announcement struct{
	AnnID int
	AnnTitle string
	AnnText string
	AnnDate string
	AnnActive bool
}
