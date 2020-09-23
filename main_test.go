package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// func TestFileExist(t *testing.T) {

// 	exist_or_no := fileExists("/home/chakavak/Desktop/go/maintest/main_test.go")
// 	if exist_or_no == false {
// 		t.Errorf("File Not Exist")
// 	}

// }

// func TestConfigLocation(t *testing.T) {
// 	locate := ConfigLocation()

// 	if _, err := os.Stat(locate); os.IsNotExist(err) {
// 		t.Errorf("file not exist or can't open")
// 	}

// }
type jsontest struct { // Struct For Open Json file gived from Post Request
	Name       string            `json:"Name"`
	Parameters map[string]string `json:"params"`
}

// type data struct {
// 	Name   string `json:"name"`
// 	Params param  `json:"params"`
// }
// type param struct {
// 	Context string `json:"context"`
// }
var filetest jsontest

func TestProfile(t *testing.T) {

	filetest.Name = "testin"
	param := make(map[string]string)
	param["context"] = "true"
	param["rfc2833-pt"] = "5555"
	param["sip-ip"] = "test"
	filetest.Parameters = param
	encode, _ := json.Marshal(filetest)

	url := "http://127.0.0.1:10000/profiles"

	req, _ := http.Post(url, "application/json", bytes.NewBuffer(encode))
	_ = req
}
