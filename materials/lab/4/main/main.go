package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"wyoassign/wyoassign"
)


func main() {
	wyoassign.InitAssignments()
	wyoassign.InitClasses()
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/api-status", wyoassign.APISTATUS).Methods("GET")
	router.HandleFunc("/assignments", wyoassign.GetAssignments).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.GetAssignment).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.DeleteAssignment).Methods("DELETE")		
	router.HandleFunc("/assignment", wyoassign.CreateAssignment).Methods("POST")	
	router.HandleFunc("/assignments/{id}", wyoassign.UpdateAssignment).Methods("PUT")

	//class data struct endpoints + homepage change
	router.HandleFunc("/", wyoassign.HandleTheHomePage).Methods("GET")
	router.HandleFunc("/classes", wyoassign.GetClasses).Methods("GET")
	router.HandleFunc("/class/{id}", wyoassign.GetClass).Methods("GET")
	router.HandleFunc("/class/{id}", wyoassign.DeleteClass).Methods("DELETE")
	router.HandleFunc("/class", wyoassign.CreateClass).Methods("POST")

	//http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}