package readfilestring

import (
	"fmt"
	"io/ioutil"
	"log"
)

func StringReader(filename string) {

	// read file
	content, err := ioutil.ReadFile(filename)

	// error check
	if err != nil {
		log.Fatal(err)
	}

	// read content as byte
	fmt.Println(content)

	// read content as string
	fmt.Println(string(content))

}
