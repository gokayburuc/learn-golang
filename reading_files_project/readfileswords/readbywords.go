package readfileswords

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func WordReader(filename string) {

	// read the file
	f, err := os.Open(filename)

	// read error check
	if err != nil {
		log.Fatal(err)
	}

	// defer close file
	defer f.Close()

	// create a scanner
	scanner := bufio.NewScanner(f)

	// scanner split into words
	scanner.Split(bufio.ScanWords)

	// show scanned items as text
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	//scanner error check
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
