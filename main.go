package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for Courses - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// Middleware, helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Building Crud API using mux")
	r := mux.NewRouter()

	// Seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "REACT", CoursePrice: 299, Author: &Author{Fullname: "PANDEY", Website: "jpc.com"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCorse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeonline Route</h1>"))
}

func getAllCorse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all COurses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one COurses")
	w.Header().Set("Content-Type", "application/json")

	// grab id From request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course foud=nd with given id")

}

// func getOneCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Get one course")
// 	w.Header().Set("Content-Type", "applicatioan/json")

// 	// grab id from request
// 	params := mux.Vars(r)

// 	// loop through courses, find matching id and return the response
// 	for _, course := range courses {
// 		if course.CourseId == params["id"] {
// 			json.NewEncoder(w).Encode(course)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode("No Course found with given id")
// 	return
// }

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Insert one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some DATA")
	}

	// wht about -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("NO DATA INSIDE JSON")
		return
	}

	// generate Unique id, string
	// Append course into courses
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop ,id, remove , add with myID

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TDO Send the response when ID is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Not found")
}
