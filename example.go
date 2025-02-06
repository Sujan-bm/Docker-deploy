package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	EmpId    int
	EmpNmane string
	Phone    int64
	DOB      string
}

// var EmpDetials Employee
var EmpDetials = []Employee{
	{101, "Alice", 974636728748, "04/05/1998"},
	{102, "Bob", 90947578748, "14/05/1898"},
	{103, "Charli", 809728748, "24/08/1998"},
	{104, "Dany", 774636728748, "09/03/2000"},
}

func getEmpDetails(c *gin.Context) {
	c.JSON(http.StatusOK, &EmpDetials)
}

func findById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("id not matched", err)
	}
	for i, v := range EmpDetials {
		if v.EmpId == id {
			c.JSON(http.StatusFound, EmpDetials[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
}

func addEmp(c *gin.Context) {
	var newEmp Employee
	c.ShouldBind(&newEmp)
	// if err != nil {
	// 	fmt.Errorf("error in binding values")
	// 	return
	// }
	EmpDetials = append(EmpDetials, newEmp)
	c.JSON(http.StatusCreated, newEmp)
}

func main() {
	router := gin.Default()
	router.GET("/getEmpDetails", getEmpDetails)
	router.GET("/findById/:id", findById)
	router.POST("/addEmp", addEmp)
	router.Run(":8080")
}
