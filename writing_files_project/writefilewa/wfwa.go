package writefilewa

import (
	"fmt"
	"log"
	"os"
)

func WriteFileWA(filename string) {

	// create file
	f, err := os.Create(filename)

	// err check
	if err != nil {
		log.Fatal(err)
	}

	// defer
	defer f.Close()

	// string into the byte
	myStr := "Camel Black"
	byteData := []byte(myStr)

	// byte writing
	_, err2 := f.Write(byteData)

	//err2 check
	if err2 != nil {
		log.Fatal(err2)
	}

	// another string
	newStr := "Camel White"
	newByteData := []byte(newStr)

	// index of ending byteData
	var idx int64 = int64(len(byteData))

	// write new data to the end of byteData
	_, err3 := f.WriteAt(newByteData, idx)

	// error check
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("Operation Done!")
}
