package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/praveent04/students_api/internal/config"
)

func main(){
	// load config 
    cfg := config.MustLoad()
    // custom logger but we are using inbuild logger here in this project
    // database setup
    // setup router
    router := http.NewServeMux()

    router.HandleFunc("GET /",func(w http.ResponseWriter,r *http.Request){
        w.Write([]byte("Welcome to student API"))
    })
    // setup server

    server := http.Server{
        Addr: cfg.HTTPServer.Addr,
        Handler: router,
    }
    
    fmt.Println("Server Started")


    err := server.ListenAndServe()
    if err != nil{
        log.Fatal("Failed to start Server")
    }

}