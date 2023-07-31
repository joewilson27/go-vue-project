package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type numbers struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type numsResponseData struct {
	Add float64 `json:"add"`
	Mul float64 `json:"mul"`
	Sub float64 `json:"sub"`
	Div float64 `json:"div"`
}

func process(numsdata numbers) numsResponseData {
	var numsres numsResponseData
	numsres.Add = numsdata.Num1 + numsdata.Num2
	numsres.Mul = numsdata.Num1 * numsdata.Num2
	numsres.Sub = numsdata.Num1 - numsdata.Num2
	numsres.Div = numsdata.Num1 / numsdata.Num2

	return numsres
}

func calc(w http.ResponseWriter, request *http.Request) {
	// we are declaring a JSON decoder to decode the incoming JSON data from
	// therequest's body.
	decoder := json.NewDecoder(request.Body)

	// declared structures
	var numsData numbers
	var numsResData numsResponseData

	decoder.Decode(&numsData)

	numsResData = process(numsData)

	fmt.Println(numsResData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(numsResData); err != nil {
		panic(err)
	}
}

func mainHold() {
	port := 8090
	http.HandleFunc("/", calc)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on port %d\n", port)
	http.ListenAndServe(addr, nil)
}
