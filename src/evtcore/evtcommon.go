package evtcore

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

func LogError(debug bool, err error) {
	if DEBUG {
		log.Printf("ERROR - %s\n", err)
	}
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

func RemoveBadChar(line io.Reader) (nl io.Reader) {
	badchars := []string{"\u0001", "\u0002", "\u0003"}

	buf := new(bytes.Buffer)
	buf.ReadFrom(line)

	for _, char := range badchars {
		rem := strings.Replace(buf.String(), char, "", -1)
		nl = strings.NewReader(rem)
	}
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
	stream = RemoveBadChar(stream)
	decoder := xml.NewDecoder(stream)

	dataattrs := make(map[string]interface{})

	for {

		t, err := decoder.Token()

		if err != nil || err == io.EOF {
			LogError(DEBUG, err)
			break
		}

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
