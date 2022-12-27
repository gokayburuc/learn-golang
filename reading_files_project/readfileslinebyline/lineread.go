package readfileslinebyline

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func LineRead(filename string) {

	//open file
	f, err := os.Open(filename)

	//open error check
	if err != nil {
		log.Fatal(err)
	}

	//defer close
	defer f.Close()

	// new scanner
	scanner := bufio.NewScanner(f)

	// show lines
	for scanner.Scan() {
		// show scan items as text
		fmt.Println(scanner.Text())
	}

	// scanner error check
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
