package evtcore

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	VERSION = "1.0"
)

var (
	DEBUG  = false
	STDOUT = false
	count  = 0
)

type FileHandler struct {
	File   *os.File
	Writer *io.Writer
}

func LogConsole(sout bool, message string) {
	if STDOUT {
		fmt.Println(message)
	}
}

func LogDebug(debug bool, message string) {
	if DEBUG {
		log.Printf("DEBUG - %s\n", message)
	}
}

func FileHandle() *FileHandler {
	return &FileHandler{}
}

func (f *FileHandler) FileOpen(path string) (err error) {
	f.File, err = os.Open(path)
	return err
}

func (f *FileHandler) FileClose() error {
	return f.File.Close()
}

func ReadStdin() (count int) {
	count = DecodeXMLandStreamJSON(os.Stdin)
	return
}

func EncodeJSONStruct(e *EventLog) {
	ret, err := json.Marshal(&e)

	if err != nil {
		log.Fatal(err)
	}

	str := string(ret)
	LogConsole(STDOUT, str)
}

func DecodeXMLandStreamJSON(stream io.Reader) (count int) {
	decoder := xml.NewDecoder(stream)

	dataattrs := make(map[string]interface{})

	for {

		t, _ := decoder.Token()

		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:

			if se.Name.Local == "Event" {
				count += 1
				var v Event
				var e EventLog

				decoder.DecodeElement(&v, &se)

				for _, data := range v.EventData.Data {
					if data.Key == "" {
						data.Key = "empty_value"
					}
					dataattrs[data.Key] = data.Value
				}

				e.System = v.System
				e.EvtData.Data = dataattrs
				EncodeJSONStruct(&e)
			}
		}
	}
	return
}
