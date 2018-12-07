package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "gin/models"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}
func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	p := Person{FirstName: firstName, LastName: lastName}
	id, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert success %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
func GetAllPersonsApi(c *gin.Context) {
	var person Person
	persons, err := person.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}
func GetPersonApi(c *gin.Context) {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	p := Person{Id: Id}
	person, err := p.GetPerson()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK,
		gin.H{"person": person})
}
func ModPersonApi(c *gin.Context) {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	p := Person{Id: Id}
	err := c.Bind(&p)
	if err != nil {
		log.Fatal(err.Error())
	}
	ra, err := p.UpdatePerson()
	if err != nil {
		log.Fatal(err.Error())
	}
	msg := fmt.Sprintf("Update person %d successful %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
func DelPersonApi(c *gin.Context) {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	p := Person{Id: Id}
	ra, err := p.DeletePerson()
	if err != nil {
		log.Fatal(err.Error())
	}
	msg := fmt.Sprintf("Delete person %d successful %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
