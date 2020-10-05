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

type Appending struct {
	Text  string `xml:",chardata"`
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type JsonFile struct { // Struct For Open Json file gived from Post Request
	Name       string            `json:"Name"`
	Parameters map[string]string `json:"params"`
}

var Params map[string]string

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
//Patch To XML File
func PatchingXml(FileName string, FILE string, w http.ResponseWriter, jsonfile JsonFile) {
	if FileName == "template.xml" {
		File, err := os.OpenFile(FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer File.Close()
	}
	TemplateFile, _ := os.Open("template.xml")
	defer TemplateFile.Close()
	xmlFile, _ := os.Open(FileName)

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	XmlByte, _ := ioutil.ReadAll(xmlFile)
	TemplateByte, _ := ioutil.ReadAll(TemplateFile)
	// we initialize our profile array
	var profile Profile
	var template Profile

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'profile' which we defined above
	xml.Unmarshal(XmlByte, &profile)
	xml.Unmarshal(TemplateByte, &template)
	Template_Length := len(template.Settings.Param)

	Profile_Length := len(profile.Settings.Param)

	for i := 0; i < Template_Length; i++ {
		for j := 0; j < Profile_Length; j++ {
			if profile.Settings.Param[j].Name == template.Settings.Param[i].Name {
				template.Settings.Param[i].Value = profile.Settings.Param[j].Value
			}
		}
	}

	template.Name = jsonfile.Name

	// counter := 0

	for key, value := range Params {

		for i := 0; i < Template_Length; i++ {
			if template.Settings.Param[i].Name == key {
				template.Settings.Param[i].Value = value
				delete(Params, key)
			}
			// else {

			// }

		}
	}

	var app Appending
	for key, value := range Params {

		app.Name = key
		app.Value = value

		template.Settings.Param = append(template.Settings.Param, app)
		delete(Params, key)

	}

	// _ = append(profile.Settings.Param, Params)
	file, _ := xml.MarshalIndent(template, "", " ")

	_ = ioutil.WriteFile(FILE, file, 0644)
	w.Write([]byte("OK"))

	// 	w.Header().Set("Content-Type", "application/xml")
	// 	w.Write(file)

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Post Request For Edit or Create New File

func CreateOrEditXml(FileName string, FILE string, w http.ResponseWriter, jsonfile JsonFile) {
	if FileName == "template.xml" {
		File, err := os.OpenFile(FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer File.Close()
	}
	TemplateFile, _ := os.Open("template.xml")
	defer TemplateFile.Close()
	xmlFile, _ := os.Open(FileName)

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	XmlByte, _ := ioutil.ReadAll(xmlFile)
	TemplateByte, _ := ioutil.ReadAll(TemplateFile)
	// we initialize our profile array
	var profile Profile
	var template Profile

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'profile' which we defined above
	xml.Unmarshal(XmlByte, &profile)
	xml.Unmarshal(TemplateByte, &template)
	Template_Length := len(template.Settings.Param)

	template.Name = jsonfile.Name

	// counter := 0

	for key, value := range Params {

		for i := 0; i < Template_Length; i++ {
			if template.Settings.Param[i].Name == key {
				template.Settings.Param[i].Value = value
				delete(Params, key)
			}
			// else {

			// }

		}
	}

	var app Appending
	for key, value := range Params {

		app.Name = key
		app.Value = value

		template.Settings.Param = append(template.Settings.Param, app)
		delete(Params, key)

	}

	// _ = append(profile.Settings.Param, Params)
	file, _ := xml.MarshalIndent(template, "", " ")

	_ = ioutil.WriteFile(FILE, file, 0644)
	w.Write([]byte("OK"))

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ReadingXML(w http.ResponseWriter, r *http.Request, Location string) {
	params, ok := r.URL.Query()["name"]

	if !ok || len(params[0]) < 1 {
		fmt.Fprintf(w, "You Should Pass One Parameter named by 'name' ")

	} else {

		// Query()["key"] will return an array of items,
		// we only want the single item.
		param := params[0]
		FileName := Location + param + ".xml"
		exist_or_no := fileExists(FileName)
		if exist_or_no == true {
			xmlFile, _ := os.Open(FileName)

			// defer the closing of our xmlFile so that we can parse it later on
			defer xmlFile.Close()

			byteValue, _ := ioutil.ReadAll(xmlFile)

			// we initialize our profile array
			var profile Profile
			// we unmarshal our byteArray which contains our
			// xmlFiles content into 'profile' which we defined above
			xml.Unmarshal(byteValue, &profile)

			file, _ := xml.MarshalIndent(profile, "", " ")

			w.Header().Set("Content-Type", "application/xml")
			w.Write(file)
		} else {
			fmt.Fprintf(w, "File not Exist")
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Give Location files from config.json

func ConfigLocation() string {
	configfile, _ := os.Open("config.json")
	byteconfig, _ := ioutil.ReadAll(configfile)
	var configstruct ConfigFile
	json.Unmarshal(byteconfig, &configstruct)
	return configstruct.Location
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// For Handling Patch Request

func PatchRequestHandling(jsonfile JsonFile, Location string, w http.ResponseWriter) {
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
		PatchingXml(FileName, FileName, w, jsonfile)

	} else {
		PatchingXml("template.xml", FileName, w, jsonfile)

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// profile for parsing xml files and give request post in json file  and pass them to WritingXML() function

func PostRequestHandling(jsonfile JsonFile, Location string, w http.ResponseWriter) {

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
		CreateOrEditXml(FileName, FileName, w, jsonfile)

	} else {
		CreateOrEditXml("template.xml", FileName, w, jsonfile)

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// profile for parsing xml files and give request post in json file  and pass them to functions
func Profiles(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Location := ConfigLocation() + "/"

	if r.Method == http.MethodGet {

		ReadingXML(w, r, Location)

	} else if r.Method == http.MethodPatch {

		// fmt.Fprintf(w, "Post Request")
		var jsonfile JsonFile

		Params = make(map[string]string)

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&jsonfile)
		if err != nil {

			panic(err)
		}

		PatchRequestHandling(jsonfile, Location, w)

	} else if r.Method == http.MethodPost {

		var jsonfile JsonFile

		Params = make(map[string]string)

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&jsonfile)
		if err != nil {

			panic(err)
		}

		PostRequestHandling(jsonfile, Location, w)

	}

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Function for Handling rouing
func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/profiles", Profiles)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("0.0.0.0:10000", nil))

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Main function
func main() {

	handleRequests()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
