package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	r.GET("/course/:id", getSingleCourse)

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
	c.Header("Content-Type", "application/json")

	// grab id from request
	id := c.Param("id")

	// loop through courses and find matching id and return the response
	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(c.Writer).Encode(course) // line for mux
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})

	// json.NewEncoder(c.Writer).Encode("No Course found with the given id")
	// return
}

func createOneCourse(c *gin.Context) {

	fmt.Println("------------Create One course-------")
	c.Header("Content-Type", "application/json")

	// what if: body is empty
	if c.Request.Body == nil {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "Please send some data",
		})
	}

	// what about {} -->> user is sending {}

	var course Course
	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	courses = append(courses, course)

	c.JSON(http.StatusCreated, course)

}
