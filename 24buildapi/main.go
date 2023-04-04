package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// model for courses -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper -- file
func IsEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

	r := gin.Default()
	r.GET("/", serveHome)
	r.GET("/all", getAllCourses)

	r.Run(":8080")
}

//controllers -- different file

//serve home route

func serveHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Home Page",
	})
	// c.Header("Content-Type", "text/html")
	// html := "<h1>Hello World!</h1>"
	// fmt.Fprint(c.Writer, html)
}

func getAllCourses(c *gin.Context) {
	fmt.Println("Get all courses")
	c.Header("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(courses)
}

func getSingleCourse(c *gin.Context) {
	fmt.Println("------------Getting One course-------")

}
