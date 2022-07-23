package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func OutputCpp(jobID string) string {
	path := fmt.Sprintf("components/"+"cpp/"+"output/"+jobID+".exe")

	cmd := exec.Command(path)
	
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
    	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	return out.String()
}

func OutputPy(jobID string) string {
	path := fmt.Sprintf("components/"+"py/"+"code/"+jobID+".py")

	cmd := exec.Command("python",path)
	
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
    	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	return out.String()
}