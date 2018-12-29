package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "gin/models"

	"github.com/gin-gonic/gin"
)

type AddPersonRequest struct {
	FirstName string `key:"first_name" required:"true"`
	LastName  string `key:"last_name" required:"true"`
}

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}
func AddPersonApi(c *gin.Context) {
	//firstName := c.Request.FormValue("first_name")
	req := new(AddPersonRequest)
	fmt.Println(c.Request.URL.Query().Get("first_name"))
	err := InputRequest(c.Request, req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(req)
	//lastName := c.Request.FormValue("last_name")
	if req.FirstName == "" || req.LastName == "" {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "nonon",
		})
	}
	p := Person{FirstName: req.FirstName, LastName: req.LastName}
	id, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert success %d", id)
	OutputResponse(c.Writer,"haha")
	fmt.Println(c.Writer.Header())
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
