package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func GenerateFile(format string, content string) string {
	
	id := uuid.New()
	filepath := fmt.Sprintf("components/%s/code/%s.%s",format,id.String(),format)

	file, err := os.Create(filepath)
	
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	file.WriteString(content)

	return filepath
	
}