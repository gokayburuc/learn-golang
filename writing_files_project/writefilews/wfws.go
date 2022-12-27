package writefilews

import (
	"fmt"
	"log"
	"os"
)

func WriteFileWS(filename string) {

	// create a new file
	f, err := os.Create(filename)

	// err check
	if err != nil {
		log.Fatal(err)
	}

	// defer
	defer f.Close()

	// write into file
	myString := "Black Hawk Down!"
	_, err2 := f.WriteString(myString)

	// err2 check
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Operation Done!")
}
