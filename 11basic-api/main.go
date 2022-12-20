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

// in sep file
type Course struct {
	CourseId    string  `jso:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB

var Courses []Course

//hleper Method

func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

//controllers in sep file

// serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveHome ")
	w.Write([]byte("<h1>Hello welcome to course API</h1>"))
}
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("NO COURSE FOUND")
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data in JSON")
		return
	}
	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(1000))

	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)

}
func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")
	index := -1

	params := mux.Vars(r)
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	for i, c := range Courses {
		if c.CourseId == params["id"] {
			index = i
			break
		}
	}
	course.CourseId = params["id"]
	if index != -1 {
		Courses = append(Courses[:index], Courses[index+1:]...)
		Courses = append(Courses, course)
		json.NewEncoder(w).Encode(course)
	} else {
		json.NewEncoder(w).Encode("enter valid course id")
	}

}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:i], Courses[i+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found")

}
func main() {
	fmt.Println("About api creation in golang")

	r := mux.NewRouter()

	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	//seeding

	Courses = append(Courses, Course{CourseId: "12", CourseName: "GoLang", CoursePrice: 499, Author: &Author{Fullname: "hemanth", Website: "LearnCodeOnline"}})
	Courses = append(Courses, Course{CourseId: "15", CourseName: "ReactJS", CoursePrice: 599, Author: &Author{Fullname: "Anvesh", Website: "LearnCodeOnline"}})
	log.Fatal(http.ListenAndServe(":4000", r))
}
