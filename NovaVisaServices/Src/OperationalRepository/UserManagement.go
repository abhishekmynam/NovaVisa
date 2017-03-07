package OperationalRepository

import (SR "SourceRepository"
	CR "ConfigRepository"
)

type UserManagement interface{
	NewUserAddition(newUser CR.User)string
	UserUpdate(newUser CR.User)string
	UserDeactivate(newUser CR.User)string
	GetAllUsers()[]CR.UserCollStruct
	AuthUser(email string, pswd string)bool
}
type userMgmt struct{}

func UserManage() UserManagement{
	return userMgmt{}
}

func (u userMgmt)NewUserAddition(newUser CR.User)string{
	var statusMsg string
	statusMsg = SR.AddToDB().AddUserToDB(newUser)
	return statusMsg
}

func (u userMgmt)UserUpdate(newUser CR.User)string{
	var statusMsg string
	newUser.Pswd =SR.FuncRep().EncryptPswd(newUser.Pswd)
	statusMsg = SR.UpdateData().UpdateUser(newUser)
	return statusMsg
}

func (u userMgmt)UserDeactivate(newUser CR.User)string{
	var statusMsg string
	newUser.IsActive = false
	statusMsg = SR.UpdateData().UpdateUser(newUser)
	return statusMsg
}

func (u userMgmt)GetAllUsers()[]CR.UserCollStruct{
	var userList []CR.UserCollStruct
	userList = SR.GetFromDB().GetUserList()
	return userList
}

func (u userMgmt)AuthUser(email string, pswd string)bool{
	pswdGot := SR.GetFromDB().GetAuthPswd(email)
	if(pswdGot == pswd){return true}
	return false
}

