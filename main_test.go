package main

import (
	"os"
	"testing"
)

func TestFileExist(t *testing.T) {

	exist_or_no := fileExists("/home/chakavak/Desktop/go/maintest/main_test.go")
	if exist_or_no == false {
		t.Errorf("File Not Exist")
	}

}

func TestConfigLocation(t *testing.T) {
	locate := ConfigLocation()

	if _, err := os.Stat(locate); os.IsNotExist(err) {
		t.Errorf("file not exist or can't open")
	}

}
