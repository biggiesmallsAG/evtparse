# evtparse
Evt(x) Parser that takes XML and outputs JSON
```
	########
	Evtparse.GO version 1.01 (8/12/2016), by Daniel Eden, SecureWorks.
	Input: XML, Output: JSON
	
	To run this program you need to have https://github.com/williballenthin/python-evtx installed,
	the scripts/evtxdump.py output is a XML stream to which you can pipe via stdin or dump to a file
	and read in to evtxparse.go.

	- Version 1.0
		> First realease of code and concept
		> Ability to stream STDIN or from file
		> Currently only able to do "System" eventlogs

	########

Usage: evtxparse [OPTIONS] argument ... 
  -d	Turn on console level debugging.
  -f string
    	Read from file. 
  -h	Display use flags.
  -o	Write JSON to stdout.
  -s	Read from stdin xml stream.
  -w string
    	Write output to file. **Placeholder if needed later**
```
