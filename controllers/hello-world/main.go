package helloworld

import (
	"encoding/json"
	"net/http"
)

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HelloWord(res http.ResponseWriter, req *http.Request) {
	response := DefaultResponse{Status: 200, Message: "Rest Api With Go"}
	jsonResponse, _ := json.Marshal(response)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(response.Status)
	res.Write(jsonResponse)
}
