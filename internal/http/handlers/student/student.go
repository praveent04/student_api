package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/praveent04/students_api/internal/types"
	"github.com/praveent04/students_api/internal/utils/response"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){

        var student types.Student

        // decoding json data into student variable

        err := json.NewDecoder(r.Body).Decode(&student)

        // checking if incoming json data is empty or not
        if errors.Is(err, io.EOF){
            response.WriteJson(w, http.StatusBadRequest,err.Error())
            return
        }

        slog.Info("creating a student")
        response.WriteJson(w,http.StatusCreated, map[string]string{"success":"OK"})
       // w.Write([]byte("Welcome to student API"))
    }
}