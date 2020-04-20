package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/vamshi/go-crud/model"
	"github.com/vamshi/go-crud/repository"
)


func userHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	
	log.Print(path)
	users := repository.GetUserData()
	// set response type as json
	w.Header().Set("Content-Type","application/json")
	//converting the users slice to json
	json.NewEncoder(w).Encode(users)
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
		Addr:         ":8080",
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



	//TODO . Will userHandler() work instead of userHandler
	r.HandleFunc("/users",userHandler)
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