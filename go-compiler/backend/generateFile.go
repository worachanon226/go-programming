package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func GenerateFile(format string, content string) string {
	
	id := uuid.New()
	path := fmt.Sprintf("components/%s/code/%s.%s",format,id.String(),format)

	file, err := os.Create(path)
	
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	file.WriteString(content)

	return ExecuteCpp(path)
}