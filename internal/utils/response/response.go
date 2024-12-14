package response

import (
	"encoding/json"
	"net/http"
)

type Response struct{
	Status string
	Error string
}

const (
	StatusOK = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

//handling the error format that will show when api will test

func GenralError (err error) Response{
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}