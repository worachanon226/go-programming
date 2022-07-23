package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func GenerateFile(format string, content string){
	
	id := uuid.New()
	path := fmt.Sprintf("./components/%s.%s",id.String(),format)

	file, err := os.Create(path)
	
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	file.WriteString(content)
}