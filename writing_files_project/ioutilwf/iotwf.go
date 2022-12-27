package ioutilwf

import (
	"fmt"
	"io/ioutil"
	"log"
)

func IotWriter() {

	// sample sentence
	val := "black hawk down!\n"

	// make byte
	data := []byte(val)

	// write into a file
	err := ioutil.WriteFile("myfile.txt", data, 0)

	// err check
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Operation Done!")
}
