package main

import (
	"os"
)

func readFile(filePath string) (string, error) {
	contend ,err:=os.ReadFile(filePath)
	if  err != nil {
		return  "" ,err
	}
	return string(contend), nil
}