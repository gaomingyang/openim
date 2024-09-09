package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type errMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func respCheck(resp *http.Response) {
	if resp.StatusCode > 299 {
		fmt.Println("Server responded with status code", resp.StatusCode)
		if resp.Header.Get("Content-Type") == "application/json" {
			var errMsg errMessage
			err := json.NewDecoder(resp.Body).Decode(&errMsg)
			errorCheck(err)
			fmt.Printf("%s: %s\n", errMsg.Status, errMsg.Message)
		}
	}
}
