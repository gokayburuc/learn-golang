package readfileswords

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// filename string
func ReadWords(filename string) {

	// read file
	f, err := os.Open(filename)

	// read err check
	if err != nil {
		log.Fatal(err)
	}

	// defer close file
	defer f.Close()

	// create a scanner
	scanner := bufio.NewScanner(f)

	// split words
	scanner.Split(bufio.ScanWords)

	// show scanned items as texts
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// scanner error check
	if err2 := scanner.Err(); err2 != nil {
		fmt.Println(err2)
		log.Fatal(err2)
	}

}
