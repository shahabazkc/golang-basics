package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to mymodules")
	r := mux.NewRouter()
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "js",
		CoursePrice: 120,
		Author:      &Author{Fullname: "shahabaz", Website: "hello.com"},
	})

	//routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCources).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to mymodules api"))
}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Println(params)

	// loop through courses and find the course with the id
	for _, item := range courses {
		if item.CourseId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// Check only if title is duplicate
	for index, item := range courses {
		if item.CourseName == course.CourseName && index >= 0 {
			json.NewEncoder(w).Encode("Course Already exist")
			return
		}
	}

	// generate a unique id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000000))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
}