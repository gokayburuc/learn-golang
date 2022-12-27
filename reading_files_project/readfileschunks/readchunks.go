package readfileschunks

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func ChunkReader(filename string) {

	// open file
	f, err := os.Open(filename)

	// open error check
	if err != nil {
		log.Fatal(err)
	}

	// defer close
	defer f.Close()

	// bufio reader
	reader := bufio.NewReader(f)

	// 16 byte array
	buf := make([]byte, 16)

	//

	for {
		// reading file into a buffer which has 16 byte size
		n, err := reader.Read(buf)

		if err != nil {
			// checking the EOF
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}

		// reading file from buffer [0:n]
		fmt.Print(string(buf[0:n]))
	}
	fmt.Println()
}
