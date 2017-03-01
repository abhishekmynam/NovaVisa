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
