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

type JsonFile struct {
	Name       string            `json:"Name"`
	Parameters map[string]string `json:"params"`
}

var jsonfile JsonFile
var Params map[string]string

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
//Write To XML File
func WritingXML(FileName string, FILE string) {
	if FileName == "test.xml" {
		File, err := os.OpenFile(FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer File.Close()
	}

	fmt.Println(FileName)
	xmlFile, _ := os.Open(FileName)
	fmt.Println("Successfully Opened test.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var profile Profile
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &profile)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	fmt.Println(len(profile.Settings.Param))
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

func Home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
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
		if fileExists(FileName) {
			WritingXML(FileName, FileName)
		} else {
			WritingXML("test.xml", FileName)
		}

		// if we os.Open returns an error then handle it

	}

}

func handleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/profiles", Home)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

func main() {

	handleRequests()
}
