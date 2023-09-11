package utils

import (
	"fmt"
	"log"
	"os"
)

// WriteFile function for writing a file.
func WriteFile(fileName string, data string) {

	f, err := os.Create("products/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f, err = os.OpenFile("products/"+fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Writing to :", "products/"+fileName)
}
