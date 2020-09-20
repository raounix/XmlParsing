package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// struct for give all xml tags for create or modify xml files
type Profile struct {
	XMLName  xml.Name `xml:"profile"`
	Text     string   `xml:",chardata"`
	Name     string   `xml:"name,attr"`
	Aliases  string   `xml:"aliases"`
	Gateways string   `xml:"gateways"`
	Domains  struct {
		Text   string `xml:",chardata"`
		Domain struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Alias string `xml:"alias,attr"`
			Parse string `xml:"parse,attr"`
		} `xml:"domain"`
	} `xml:"domains"`
	Settings struct {
		Text  string `xml:",chardata"`
		Param []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"param"`
	} `xml:"settings"`
}

type ConfigFile struct { // Struct For Read location files from config.json
	Location string `json:"location"`
}

type JsonFile struct { // Struct For Open Json file gived from Post Request
	Name       string            `json:"Name"`
	Parameters map[string]string `json:"params"`
}

var jsonfile JsonFile        // for opening json file
var Params map[string]string // give json file parameters for search in xml files

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// If File Exist or No

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Write To XML File
func WritingXML(FileName string, FILE string) {
	if FileName == "test.xml" {
		File, err := os.OpenFile(FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer File.Close()
	}

	xmlFile, _ := os.Open(FileName)

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our profile array
	var profile Profile
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'profile' which we defined above
	xml.Unmarshal(byteValue, &profile)

	Profile_Length := len(profile.Settings.Param)
	profile.Name = jsonfile.Name
	for key, value := range Params {

		for i := 0; i < Profile_Length; i++ {
			if profile.Settings.Param[i].Name == key {
				profile.Settings.Param[i].Value = value
			}
		}
	}
	file, _ := xml.MarshalIndent(profile, "", " ")

	_ = ioutil.WriteFile(FILE, file, 0644)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Give Location files from config.json

func ConfigLocation() string {
	configfile, _ := os.Open("config.json")
	byteconfig, _ := ioutil.ReadAll(configfile)
	defer configfile.Close()
	var configstruct ConfigFile
	json.Unmarshal(byteconfig, &configstruct)
	return configstruct.Location
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// profile for parsing xml files and give request post in json file  and pass them to WritingXML() function
func Profiles(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Location := ConfigLocation() + "/"

	if r.Method == http.MethodGet {

		// f, err := os.Open("/tmp/out.xml")
		// if err != nil {
		// 	panic(err)
		// }

		// x := r.Form.Get("test")

	} else if r.Method == http.MethodPost {

		fmt.Fprintf(w, "Post Request")

		Params = make(map[string]string)

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&jsonfile)
		if err != nil {

			panic(err)
		}
		for key, value := range jsonfile.Parameters {
			Params[key] = value

		}

		var FileName = jsonfile.Name + ".xml"

		// File, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// _ = File
		FileName = Location + FileName
		if fileExists(FileName) {
			WritingXML(FileName, FileName)
		} else {
			WritingXML("test.xml", FileName)
		}

	}

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Function for Handling rouing
func handleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/profiles", Profiles)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Main function
func main() {

	handleRequests()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
