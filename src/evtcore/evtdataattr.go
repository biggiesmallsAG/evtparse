package evtcore

type EventData struct {
	Data []DataAttrs `xml:"Data"`
}

type DataAttrs struct {
	Key   string `xml:"Name,attr" json:",omitempty"`
	Value string `xml:",chardata" json:",omitempty"`
}

type DataCompiled struct {
	Data map[string]interface{}
}
