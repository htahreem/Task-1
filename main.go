package main

import (
	"errors"
	"net/http"
	"strings"

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

func studentExists(id string) bool {
	for _, val := range students {
		if strings.EqualFold(val.ID, id) {
			return true
		}
	}
	return false
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

	if studentExists(newStudent.ID) {
		context.IndentedJSON(http.StatusConflict, gin.H{"error": "Student with the same ID already exists"})
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

	if newStudent.ID != id && studentExists(newStudent.ID) {
		context.IndentedJSON(http.StatusConflict, gin.H{"error": "Student with the updated ID already exists"})
		return
	}

	currStudent.Name = newStudent.Name
	currStudent.RollNo = newStudent.RollNo
	currStudent.ContactNo = newStudent.ContactNo
	currStudent.Email = newStudent.Email

	context.IndentedJSON(http.StatusOK, currStudent)
}

func deleteStudent(context *gin.Context) {
	id := context.Param("id")
	for ind, val := range students {
		if val.ID == id {
			students = append(students[:ind], students[ind+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "Student successfully deleted"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"Error": "Student not found"})
}

func main() {
	router := gin.Default()
	router.GET("/getStudents", getStudents)
	router.GET("/getStudent/:id", getStudent)
	router.POST("/addStudent", addStudent)
	router.PUT("/updateStudent/:id", updateStudent)
	router.DELETE("/deleteStudent/:id", deleteStudent)

	router.Run("localhost:9090")
}
