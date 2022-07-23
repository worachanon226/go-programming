package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecutePy(jobID string) string {
	path := fmt.Sprintf("components/"+"py/"+"code/"+jobID+".py")

	cmd := exec.Command("python",path)
	
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
    	return fmt.Sprint(err) + ": " + stderr.String()
	}else{
	return out.String()
	}
}