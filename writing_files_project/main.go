package main

import (
	"writingfiles/formattedwrite"
	"writingfiles/ioutilwf"
	"writingfiles/writefilews"
)

func main() {
	writefilews.WriteFileWS("blackhawk.txt")
	ioutilwf.IotWriter()
	formattedwrite.FormattedWrite()

	TestFunctionTester()

}
