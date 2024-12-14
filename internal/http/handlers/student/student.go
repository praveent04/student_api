package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/praveent04/students_api/internal/types"
	"github.com/praveent04/students_api/internal/utils/response"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){

        
        slog.Info("creating a student")
        var student types.Student

        // decoding json data into student variable

        err := json.NewDecoder(r.Body).Decode(&student)

        // checking if incoming json data is empty or not
        if errors.Is(err, io.EOF){
            response.WriteJson(w, http.StatusBadRequest,response.GenralError(fmt.Errorf("empty body")))  // sending a custom error message 
            return
        }

        if err !=nil{
            response.WriteJson(w,http.StatusBadRequest,response.GenralError(err))
            return
        }

        // request validation -> we can do it manually but we are using a package here for that ( playground validator )
        // we will add extra validate argument in type.go file

        if err := validator.New().Struct(student); err != nil{
            validateErrs := err.(validator.ValidationErrors) // typecasting the normal error into validator error to pass in validate error function below

            response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrs))
        }

        response.WriteJson(w,http.StatusCreated, map[string]string{"success":"OK"})
       // w.Write([]byte("Welcome to student API"))
    }
}