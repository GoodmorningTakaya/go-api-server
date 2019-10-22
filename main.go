package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Person Object
type Person struct {
	Name       string `json: "name"`
	Age        int    `json: "age"`
	Profession string `json: "profession"`
	HairColor  string `json: "hairColor"`
}

//Handles http requests
func main() {
	http.HandleFunc("/people", postPersonOrGetPeople)
	http.HandleFunc("/people/", getByName)
	http.ListenAndServe(":8080", nil)

}

//Global map of people
var peopleMap = make(map[string]Person)

func postPersonOrGetPeople(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//Extract the raw contents of the request body
		decoder := json.NewDecoder(r.Body)
		var p Person
		err := decoder.Decode(&p)
		if err != nil {
			panic(err)
		}
		//Append the person
		peopleMap[p.Name] = p
		//Marshal the map of people
		file, err := json.MarshalIndent(peopleMap, "", "\t")
		if err != nil {
			fmt.Fprintf(w, "%s", "JSON Marshalling Error")
			w.WriteHeader(http.StatusInternalServerError)
		}
		//Write the data out to a file
		err = ioutil.WriteFile("data.json", file, 0644)
		if err != nil {
			fmt.Fprintf(w, "%s", "WriteFile Error")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "GET":
		//Open the file containing data about people
		file, err := os.Open("data.json")

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		//Create a file scanner
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		//Scan the lines into the array of strings
		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		file.Close()

		//Display the file contents onto the page
		for _, eachline := range txtlines {
			fmt.Fprintf(w, "%s\n", eachline)
		}

	}

}

func getByName(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[8:]
	val, err := json.Marshal(peopleMap[name])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "%s", val)
}
