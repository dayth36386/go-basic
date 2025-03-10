package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var DataList []Data

func init() {
	DataTest := `[{ "Id": 101, "Name": "John Doe", "Age": 20 }, { "Id": 202, "Name": "Jane Doe", "Age": 25 }]`
	err := json.Unmarshal([]byte(DataTest), &DataList)
	if err != nil {
		log.Fatal(err)
	}
}
func dataListHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		DataJson, err := json.Marshal(DataList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(DataJson)
	case http.MethodPost:
		var data Data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data.Id = len(DataList) + 1
		DataList = append(DataList, data)
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/data", dataListHandler)
	http.ListenAndServe(":8080", nil)
}
