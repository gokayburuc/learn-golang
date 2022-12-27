package readfileslinebyline

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func LineReader(filename string) {

	// open file
	f, err := os.Open(filename)

	// err check
	if err != nil {
		log.Fatal(err)
	}

	// defer close
	defer f.Close()

	//create a scanner
	scanner := bufio.NewScanner(f)

	// scanner show text
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// scanner error check
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
