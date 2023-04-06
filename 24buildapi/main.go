package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

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
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	r := gin.Default()
	r.GET("/", serveHome)
	r.GET("/course/all", getAllCourses)
	r.GET("/course/:id", getSingleCourse)
	r.POST(("/course/add"), createOneCourse)

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
	// if c.Request.Body == nil {
	// 	fmt.Println("please send some data")
	// 	c.JSON(http.StatusNoContent, gin.H{
	// 		"error": "Please send some data",
	// 	})
	// }

	// what about {} -->> user is sending {}

	var course Course
	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// checking isEmpty
	if course.IsEmpty() {
		c.JSON(http.StatusOK, gin.H{
			"error": "course is Empty. Please send some data",
		})
		return
	}

	// generate a unique id, convert into string
	// append course into course
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	c.JSON(http.StatusCreated, course)

}
