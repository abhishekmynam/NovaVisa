package OperationalRepository

import (SR "SourceRepository"
	CR "ConfigRepository"
)

type AnnouncementManagement interface{
	NewAnnouncementAdd(newAnnouncement CR.Announcement)string
	AnnouncementUpdate(newAnnouncement CR.Announcement)string
	GetActiveAnnouncements()[]CR.Announcement
	GetAllAnnouncements()[]CR.Announcement
}
type announcementMgmt struct{}

func AnnouncementManage() AnnouncementManagement{
	return announcementMgmt {}
}

func (a announcementMgmt)NewAnnouncementAdd(newAnnouncement CR.Announcement)string{
	var statusMsg string
	statusMsg = SR.AddToDB().CreateAnnouncements(newAnnouncement)
	return statusMsg
}

func (a announcementMgmt)AnnouncementUpdate(newAnnouncement CR.Announcement)string{
	var statusMsg string
	statusMsg = SR.UpdateData().UpdateAnnouncement(newAnnouncement)
	return statusMsg
}
func (a announcementMgmt)GetActiveAnnouncements()[]CR.Announcement{
	var annouceList []CR.Announcement
	annouceList = SR.GetFromDB().GetAnnouncementList()
	return annouceList
}

func (a announcementMgmt)GetAllAnnouncements()[]CR.Announcement{
	var annouceList []CR.Announcement
	annouceList = SR.GetFromDB().GetAllAnnouncementList()
	return annouceList
}