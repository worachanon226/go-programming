package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"strings"
)

func ExecuteCpp(file string) string {
	jobID := strings.Split(path.Base(file), ".")[0]
	
	outPath := path.Join("components","cpp","output",jobID)
	cmd := exec.Command("g++",file,"-o",outPath)
	
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
    	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	return jobID
}