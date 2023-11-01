package handlers

import (
	"errors"
	"example/go-api-with-db-course/database"
	"example/go-api-with-db-course/internal/api/requests"
	"example/go-api-with-db-course/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GetFactByID		godoc
//
//	@Summary		GetFactByID
//	@Description	Fetch fact data by ID.
//	@Produce		application/json
//	@Param			id	path	string	true	"Fact ID"
//	@Failure		400	"ID is not a valid number"
//	@Failure		404	"Data not found"
//	@Tags			facts
//	@Router			/facts/{id} [get]
func GetFactByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is not a valid number"})
		return
	}
	var fact models.Fact
	result := database.DB.Db.First(&fact, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, fact)
}

// ListFacts		godoc
//
//	@Summary		ListFacts
//	@Description	Fetch all facts data.
//	@Produce		application/json
//	@Param			question	query	string	false	"Search question"
//	@Tags			facts
//	@Router			/facts [get]
func ListFacts(c *gin.Context) {
	var facts []models.Fact
	question, exist := c.GetQuery("question")
	query := database.DB.Db
	if exist {
		query = query.Where("LOWER(question) LIKE LOWER(?)", "%"+question+"%")
	}

	query.Find(&facts)

	c.IndentedJSON(http.StatusOK, facts)
}

// CreateFacts		godoc
//
//	@Summary		CreateFacts
//	@Description	Create new fact data.
//	@Accept			json
//	@Produce		application/json
//	@Param			request	body		requests.CreateFactPayload	true	"Fact object to create"
//	@Failure		400		"Error parsing body payload"
//	@Tags			facts
//	@Router			/facts [post]
func CreateFacts(c *gin.Context) {
	createFactPayload := requests.CreateFactPayload{}
	if err := c.BindJSON(&createFactPayload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error parsing body payload"})
		return
	}

	fact := models.Fact{
		Question: createFactPayload.Question,
		Answer:   createFactPayload.Answer,
	}

	query := database.DB.Db.Begin()

	query.Create(&fact)

	query.Commit()

	c.IndentedJSON(http.StatusCreated, fact)
}
