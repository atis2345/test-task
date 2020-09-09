package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const PORT = 10000
const PortsApiUrl = "http://port-service:10001"
const PortsFile = "ports.json"

type JsonRecord struct {
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

func importEndpoint(w http.ResponseWriter, r *http.Request) {
	// TODO only for POST request
	// FIXME optimize for large files. Going to use json.Decoder
	f, _ := os.Open(PortsFile)
	res, _ := ioutil.ReadAll(f)

	var records map[string]JsonRecord
	json.Unmarshal(res, &records)

	request, _ := json.Marshal(records)
	b := bytes.NewBuffer(request)

	serviceResponse, _ := http.Post(PortsApiUrl+"/port/", "json", b)
	w.Header().Add("X-Total-Count", serviceResponse.Header.Get("X-Total-Count"))
}

func portEndpoint(w http.ResponseWriter, r *http.Request) {
	// TODO only for GET requests
	serviceResponse, _ := http.Get(PortsApiUrl + "/port/")
	io.Copy(w, serviceResponse.Body)
	w.Header().Add("X-Total-Count", serviceResponse.Header.Get("X-Total-Count"))
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/import/", importEndpoint)
	http.HandleFunc("/port/", portEndpoint)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v%v", ":", PORT), nil))
}

func main() {
	handleRequests()
}
