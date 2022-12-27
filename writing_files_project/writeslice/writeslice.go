package writeslice

import (
	"fmt"
	"log"
	"os"
)

func WriteSlice() {
	// create a file
	f, err := os.Create("data.txt")

	// err check
	if err != nil {
		log.Fatal(err)
	}

	// defer close file
	defer f.Close()

	// slice
	words := []string{"sky", "falcon", "hawk", "river"}

	// for range - index, slice
	for _, word := range words {
		_, err := f.WriteString(word + "\n")

		// err check
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Println("Operation Done!")
}
