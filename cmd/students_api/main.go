package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/praveent04/students_api/internal/config"
	"github.com/praveent04/students_api/internal/http/handlers/student"
)

func main(){
	// load config 
    cfg := config.MustLoad()

    // custom logger but we are using inbuild logger here in this project
    // database setup
    // setup router

    router := http.NewServeMux()

    router.HandleFunc("POST /api/students",student.New())
    // setup server

    server := http.Server{
        Addr: cfg.HTTPServer.Addr,
        Handler: router,
    }
    
    slog.Info("Server started", slog.String("address", cfg.HTTPServer.Addr))
  

    // writting a seperate go routine so that program will intrupt gracefully
    // this go routine will run concurrently

    // to make the run of server listen go routine synchronyse we wil made a seperate channel

    done := make(chan os.Signal,1)

    //this code line will trigger our signal that program is intrupt

    signal.Notify(done, os.Interrupt,syscall.SIGINT, syscall.SIGTERM)
    go func (){
        
        err := server.ListenAndServe()
        if err != nil{
            log.Fatal("Failed to start Server")
        }
    
    } ()

    // until any signel will go inside done channel our program will remain block at this point and all simultaneous routines will keep running

    <-done

    // structured log
    slog.Info("Shutting down the server")

    // writting this line so that if our server will take long time or infinite time to shut down then a context will tell us to avoid port blocking by the application 
    //context.background() is just a kind of starting pointof the timeout

    ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil{
        slog.Error("Failed to shutdown server", slog.String("error",err.Error()))
    }

    slog.Info("Server shut down successfully :)")
}