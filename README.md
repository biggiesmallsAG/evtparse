# evtparse
Evt(x) Parser that takes XML and outputs JSON
```
	########
	Evtparse.GO version 1.01 (8/12/2016), by Daniel Eden, SecureWorks.
	Input: XML, Output: JSON

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
