package evtcore

type Event struct {
	System System `xml:"System"`
	EventData EventData `xml:"EventData"`
}

type System struct {
	TimeCreated TC `xml:"TimeCreated"`
	EventID string `xml:"EventID"`
	Channel string `xml:"Channel"`
	EventRecordID string `xml:"EventRecordID"`
	Computer string `xml:"Computer"`
	Security Sec `xml:"Security"`
	Execution SubExec `xml:"Execution"`
}

type SubExec struct {
	ProcessID string `xml:"ProcessID,attr"`
	ThreadID string `xml:"ThreadID,attr"`
}

type Sec struct {
	UserID string `xml:"UserID,attr"`
}

type TC struct {
	SystemTime string `xml:"SystemTime,attr"`
}
