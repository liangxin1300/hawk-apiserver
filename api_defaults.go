package main

import (
	"encoding/xml"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func handleApiRscDefaults(w http.ResponseWriter, r *http.Request, cib_data string) bool {
	// parse xml into Cib struct
	var cib Cib
	err := xml.Unmarshal([]byte(cib_data), &cib)
	if err != nil {
		log.Error(err)
		return false
	}

	cib.Configuration.URLType = "rsc_defaults"

	w.Header().Set("Content-Type", "application/json")

	jsonData, jsonError := MarshalOut(r, &cib)
	if jsonError != nil {
		log.Error(jsonError)
		return false
	}

	io.WriteString(w, string(jsonData)+"\n")
	return true
}

func handleApiOpDefaults(w http.ResponseWriter, r *http.Request, cib_data string) bool {
	// parse xml into Cib struct
	var cib Cib
	err := xml.Unmarshal([]byte(cib_data), &cib)
	if err != nil {
		log.Error(err)
		return false
	}

	cib.Configuration.URLType = "op_defaults"

	w.Header().Set("Content-Type", "application/json")

	jsonData, jsonError := MarshalOut(r, &cib)
	if jsonError != nil {
		log.Error(jsonError)
		return false
	}

	io.WriteString(w, string(jsonData)+"\n")
	return true
}