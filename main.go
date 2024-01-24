package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	RollNo    int    `json:"rollno"`
	ContactNo int    `json:"contactno"`
	Email     string `json:"email"`
}

var students = []Student{
	{ID: "1", Name: "John", RollNo: 1, ContactNo: 1234, Email: "john@gmail.com"},
	{ID: "2", Name: "Alice", RollNo: 2, ContactNo: 2245, Email: "alice@gmail.com"},
	{ID: "3", Name: "Bob", RollNo: 3, ContactNo: 3566, Email: "bob@gmail.com"},
	{ID: "4", Name: "Gwen", RollNo: 4, ContactNo: 7654, Email: "gwen@gmail.com"},
}

func getStudents(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, students)
}

func addStudent(context *gin.Context) {
	var newStudent Student
	if err := context.BindJSON(&newStudent); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	students = append(students, newStudent)
	context.IndentedJSON(http.StatusCreated, students)
}

func getStudentByID(id string) (*Student, error) {
	for ind, val := range students {
		if val.ID == id {
			return &students[ind], nil
		}
	}
	return nil, errors.New("Student doesn't exist")
}

func getStudent(context *gin.Context) {
	id := context.Param("id")
	student, err := getStudentByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, student)
}

func updateStudent(context *gin.Context) {
	var newStudent Student
	if err := context.BindJSON(&newStudent); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := context.Param("id")
	currStudent, err := getStudentByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	currStudent.Name = newStudent.Name
	currStudent.RollNo = newStudent.RollNo
	currStudent.ContactNo = newStudent.ContactNo
	currStudent.Email = newStudent.Email

	context.IndentedJSON(http.StatusOK, currStudent)
}

func main() {
	router := gin.Default()
	router.GET("/getStudents", getStudents)
	router.GET("/getStudent/:id", getStudent)
	router.POST("/addStudent", addStudent)
	router.PATCH("/updateStudent/:id", updateStudent)
	router.Run("localhost:9090")
}
