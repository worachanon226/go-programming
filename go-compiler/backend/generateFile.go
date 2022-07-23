package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

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

	if format=="cpp" {
		return ExecuteCpp(filepath)
	}else{
		return strings.Split(path.Base(filepath), ".")[0]
	}
	
}