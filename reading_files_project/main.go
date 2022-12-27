package main

import (
	"readfiles/readfileschunks"
)

func main() {
	// readfilestring.StringReader("test.txt")
	// readfileslinebyline.LineReader("test.txt")

	// readfileslinebyline.LineRead("test.txt")

	// readfileswords.WordReader("test.txt")

	// readfileschunks.ChunkReader("test.txt")

	// readfileswords.ReadWords("test.txt")
	readfileschunks.ChunkReader("test.txt")
    
    
}
