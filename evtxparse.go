/*
	EVTXParse.GO by Daniel Eden, SecureWorks.
*/

package main

import (
	"evtcore"
	"flag"
	"fmt"
	"log"
	"os"
)

type RuntimeOptions struct {
	Debug, Stdin, Help, Stdout bool
	WriteOutput, ReadFile      string
}

func fUsage() {
	fmt.Println("\n\t########")
	fmt.Printf("\tEvtparse.GO version %s (5/12/2016), by Daniel Eden, SecureWorks.\n\tInput: XML, Output: JSON\n", evtcore.VERSION)
	fmt.Println("\t########\n")
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Usage = fUsage

	var count = 0
	var runopt RuntimeOptions

	flag.BoolVar(&runopt.Help, "h", false, "Display use flags.")
	flag.BoolVar(&runopt.Debug, "d", false, "Turn on console level debugging.")
	flag.BoolVar(&runopt.Stdin, "s", false, "Read from stdin xml stream.")
	flag.BoolVar(&runopt.Stdout, "o", false, "Write JSON to stdout.")
	flag.StringVar(&runopt.ReadFile, "f", "", "Read from file.")
	flag.StringVar(&runopt.WriteOutput, "w", "", "Write output to file. **Placeholder if needed later**")

	flag.Parse()

	if runopt.Help || !runopt.Stdin && runopt.ReadFile == "" {
		fUsage()
	}

	if runopt.Debug {
		evtcore.DEBUG = true
	}

	if runopt.Stdout {
		evtcore.STDOUT = true
	}

	if runopt.Stdin {
		evtcore.LogDebug(evtcore.DEBUG, "Reading stream from stdin..")

		count = evtcore.ReadStdin()

		msg := fmt.Sprintf("Event Count: %d", count)
		evtcore.LogDebug(evtcore.DEBUG, msg)
	}

	if runopt.ReadFile != "" {
		fH := evtcore.FileHandle()
		err := fH.FileOpen(runopt.ReadFile)

		if err != nil {
			log.Fatal(err)
		}

		msg := fmt.Sprintf("%s opened for reading..", runopt.ReadFile)
		evtcore.LogDebug(evtcore.DEBUG, msg)

		defer fH.FileClose()

		count = evtcore.DecodeXMLandStreamJSON(fH.File)

		msg = fmt.Sprintf("Event Count: %d", count)
		evtcore.LogDebug(evtcore.DEBUG, msg)
	}
}
