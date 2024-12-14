package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct{
	Status string `json:"status"`
	Error string `json:"error"`
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

func ValidationError(errs validator.ValidationErrors) Response{
	var errMsgs []string
	for _,err := range errs{
		
		// we get actual tag a function in validator package to read validator fields in incoming err messages

		switch err.ActualTag(){ 
	    case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is Invalid", err.Field()))	
		}
	}
	return Response{
		Status: StatusError,
		Error: strings.Join(errMsgs,","),
	}
}