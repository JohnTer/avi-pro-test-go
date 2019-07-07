package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseID struct {
	Data string `json:"data"`
	Err  int    `json:"err"`
}

func SendResponse(w *http.ResponseWriter, id string, err int) {
	resp := &ResponseID{Data: id, Err: err}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(200)
	js, qqqq := json.Marshal(resp)
	_ = qqqq
	(*w).Write(js)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	id := keys[0]
	randStr, err := GetModel(id)

	SendResponse(&w, randStr, err)

}

func postHandler(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		SendResponse(&w, "", 1)
	}
	id, ierr := PostModel(string(b))
	if ierr != 0 {
		SendResponse(&w, "", 1)
	} else {
		SendResponse(&w, id, ierr)
	}
}
