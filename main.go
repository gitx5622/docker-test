package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gits/docker-test/model"
	"github.com/gits/docker-test/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

)


func userHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	
	log.Print(path)
	vars := mux.Vars(req)
	// retreive the id from the input request
	id := vars["id"]

	i,err := strconv.Atoi(id)

	log.Print(id)
	if err != nil {
		fmt.Printf("unable to convert id:%v to int",id)
	}
	user := repository.GetUserData(i)
	// set response type as json
	w.Header().Set("Content-Type","application/json")
	//converting the users slice to json
	json.NewEncoder(w).Encode(user)
}

func addUserHandler(w http.ResponseWriter,req *http.Request){
	path := req.URL.Path
	log.Printf("path is %v",path)

	decoder := json.NewDecoder(req.Body)

	var user model.User
	// decode the data and assign it to model.User type
	err := decoder.Decode(&user)

	if err != nil {
		log.Fatalf("unable to decode body %v",err)
	}

	repository.InsertRecord(user)
	log.Printf("Inserted record is %+v",user)

}
func main() {
		// Create Server and Route Handlers
		r := mux.NewRouter()
		
	srv := &http.Server{
		Handler:      r,
		Addr:         ":9000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := repository.MigrateDb()

	if err != nil {
		log.Fatalf("Unable to perform db migration %v",err)
	}
	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()



	//TODO . Will userHandler() with brace. work instead of userHandler
	r.HandleFunc("/users/{id:[0-9]+}",userHandler)
	r.HandleFunc("/addUser", addUserHandler)

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}