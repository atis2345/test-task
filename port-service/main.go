package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const PORT = 10001

var database = make(map[string]PortRecord)

type PortRecord struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`   // FIXME unknown type
	Regions     []string  `json:"regions"` // FIXME unknown type
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func portsEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "PUT" {
		var records map[string]PortRecord
		body, _ := ioutil.ReadAll(r.Body)

		json.Unmarshal(body, &records)

		for k, v := range records {
			database[k] = v
		}

	} else if r.Method == "GET" {
		response, _ := json.Marshal(database)
		w.Write(response)
	}

	w.Header().Add("X-Total-Count", strconv.Itoa(len(database)))
	w.WriteHeader(200)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/port/", portsEndpoint)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v%v", ":", PORT), nil))
}

func main() {
	handleRequests()
}
