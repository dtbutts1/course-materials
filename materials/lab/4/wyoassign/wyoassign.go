package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	Assignments []Assignment `json:"assignments"`
}

//added response struct
type MyResponse struct{
	Classes []Class `json:"classes"`
}

type Assignment struct {
	Id string `json:"id"`
	Title string `json:"title`
	Description string `json:"desc"`
	Points int `json:"points"`
}

//added class struct
type Class struct{
	Id string `json:"id"`
	CourseNumber int `json:"coursenumber`
	Name string `json:"name"`
	Professor string `json:"professor"`
	Department string `json:"department"`
}

var Assignments []Assignment
const Valkey string = "FooKey"

var Classes []Class
const Valkey2 string = "BarKey"

func InitAssignments(){
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	Assignments = append(Assignments, assignmnet)
}

func InitClasses(){
	var class Class
	class.Id = "Cyber1"
	class.CourseNumber = 1100
	class.Department = "COSC"
	class.Name = "Topics in Cybersecurity"
	class.Professor = "Mike Borowzek"		//spelled that wrong haha
	Classes = append(Classes, class)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func HandleTheHomePage(w http.ResponseWriter, r *http.Request){
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You Have Made it to the HomePage BabyGirl")
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

//added get classes function
func GetClasses(w http.ResponseWriter, r *http.Request){
	log.Printf("Entering %s end point", r.URL.Path)
	var response MyResponse
	response.Classes = Classes

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

		jsonResponse, err := json.Marshal(response)
	if err!=nil{
		return
	}
	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)
	response["status"] = "No assignments associated with this endpoint"

	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			response["status"] = "Worked"
			json.NewEncoder(w).Encode(assignment)
			break
		}
	}

	//TODO : Provide a response if there is no such assignment
	if(response["status"] != "Worked"){
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return 
		}
		w.Write(jsonResponse)
	}
}

//added get class function. Returns one specific class based off course number
func GetClass(w http.ResponseWriter, r *http.Request){
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)
	response["status"] = "No classes associated with this endpoint"

	for _, class := range Classes {
		if class.Id == params["id"]{
			response["status"] = "Worked"
			json.NewEncoder(w).Encode(class)
			break
		}
	}

	if(response["status"] != "Worked"){
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return 
		}
		w.Write(jsonResponse)
	}
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

//added delete class function
func DeleteClass(w http.ResponseWriter, r *http.Request){
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such coursenumber to Delete"
	for index, class := range Classes {
		if class.Id == params["id"]{
			Classes = append(Classes[:index], Classes[index+1:]...)
			response["status"] = "Success"
			break
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	messageResponse := make(map[string]string)
	messageResponse["status"] = "Failed"
	
	var response Response
	response.Assignments = Assignments
	
	for index, assignment := range response.Assignments {
		if assignment.Id == r.FormValue("id") {
			messageResponse["status"] = "Worked"	//signal that it worked

			Assignments = append(Assignments[:index], Assignments[index+1:]...)
			var newPost Assignment
			newPost.Id =  r.FormValue("id")
			newPost.Title =  r.FormValue("title")
			newPost.Description =  r.FormValue("desc")
			newPost.Points, _ =  strconv.Atoi(r.FormValue("points"))
			Assignments = append(Assignments, newPost)

			w.WriteHeader(http.StatusCreated)
		}
	}

	//didn't find assignment to udpate
	if(messageResponse["status"] != "Worked"){
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return 
		}
		w.Write(jsonResponse)
	}else{
		//found assignment to change
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return 
		}
		w.Write(jsonResponse)
	}


}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignmnet Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		assignmnet.Id =  r.FormValue("id")
		assignmnet.Title =  r.FormValue("title")
		assignmnet.Description =  r.FormValue("desc")
		assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
		Assignments = append(Assignments, assignmnet)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)

}

//added create class function
func CreateClass(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var class Class
	r.ParseForm()
	if(r.FormValue("id") != ""){
		class.Id = r.FormValue("id")
		class.Name = r.FormValue("name")
		class.CourseNumber, _ = strconv.Atoi(r.FormValue("coursenumber"))
		class.Department = r.FormValue("department")
		class.Professor = r.FormValue("professor")
		Classes = append(Classes, class)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)
}