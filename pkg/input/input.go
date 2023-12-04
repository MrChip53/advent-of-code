package input

import (
	"bufio"
	"log"
	"os"
)

func OpenInputText(filePath string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return file, bufio.NewScanner(file)
}
