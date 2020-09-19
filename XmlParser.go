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

type Person struct {
	Name       string   `json:"Name"`
	Parameters []string `json:"parameters"`
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
		var Params map[string]string

		Params = make(map[string]string)
		for key, value := range r.Form {
			Params[key] = value[0]

		}
		decoder := json.NewDecoder(r.Body)
		var t Person

		err := decoder.Decode(&t)
		if err != nil {

			panic(err)
		}
		log.Println(t.Parameters)

		var FileName = t.Name + ".xml"

		File, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_ = File
		xmlFile, err := os.Open("test.xml")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}

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
		Profile_Length := len(profile.Settings.Param)
		profile.Name = t.Name
		for key, value := range Params {

			for i := 0; i < Profile_Length; i++ {
				if profile.Settings.Param[i].Name == key {
					profile.Settings.Param[i].Value = value
				}
			}
		}
		file, _ := xml.MarshalIndent(profile, "", " ")

		_ = ioutil.WriteFile(FileName, file, 0644)
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
