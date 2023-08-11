package main

import (
	"net/http"
	"strconv"

	"automation-challenge/internal/model"
	"github.com/gin-gonic/gin"
)

type routerDependencies struct {
	model model.PackageInterface
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("internal/templates/*")
	rd := routerDependencies{model: model.NewPackage()}
	// Views:
	r.GET("/world", rd.getWorldView)
	r.GET("/persons", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create-person.html", gin.H{
			"Countries": rd.model.GetCountries(),
		})
	})

	// BFF:
	r.GET("/bff/persons", func(c *gin.Context) {
		c.JSON(http.StatusOK, rd.model.GetPeople())
	})
	r.GET("/bff/countries", func(c *gin.Context) {
		c.JSON(http.StatusOK, rd.model.GetCountries())
	})
	r.POST("/bff/persons", rd.createPerson)
	r.DELETE("/bff/persons/:id", rd.deletePerson)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (rd *routerDependencies) getWorldView(c *gin.Context) {
	gojsModel, err := rd.model.GetWorldViewModel()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "world.html", err.Error())
		return
	}
	c.HTML(http.StatusOK, "world.html", gin.H{
		"model": gojsModel,
	})
}

func (rd *routerDependencies) createPerson(c *gin.Context) {
	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := rd.model.CreatePerson(person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (rd *routerDependencies) deletePerson(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = rd.model.DeletePerson(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Person deleted successfully"})
}
