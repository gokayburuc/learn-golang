package formattedwrite

import (
	"fmt"
	"log"
	"os"
)

func FormattedWrite() {
	// create a file
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	// defer close
	defer f.Close()

	// define two constants
	const name, age = "Bruce Wayne", 40

	//write into file f formatted text
	n, err := fmt.Fprintln(f, name, "is", age, "years old.")

	// err check
	if err != nil {
		log.Fatal(err)
	}

	//  end message
	fmt.Println(n, "bytes written")
	fmt.Println("Operation Done!")
}
