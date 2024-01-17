package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	models "github.com/italobbarros/go-mqtt-broker/internal/api/models"
)

// Endpoints for Publish

// createPublish cria uma nova publicação.
// @Summary Create a new publish
// @Description Create a new publish
// @Tags Publishs
// @Accept json
// @Produce json
// @Param input body models.PublishRequest true "Publish object that needs to be added"
// @Success 201 {object} models.PublishRequest
// @Router /publishs [post]
func (r *Routes) CreatePublish(c *gin.Context) {
	var publish models.Publish
	if err := c.ShouldBindJSON(&publish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publish.Timestamp = time.Now()
	publish.NumberTimestamp = time.Now().Unix()
	r.db.Create(&publish)

	//Update Topic
	var existingTopic models.Topic

	if err := r.db.Where("\"Name\" = ?", publish.TopicName).First(&existingTopic).Error; err != nil {
		// Tópico não existe, então criamos um novo
		var newTopic models.Topic
		newTopic.Name = publish.TopicName
		newTopic.Publish = publish
		newTopic.Created = time.Now()
		newTopic.Updated = time.Now()
		r.db.Create(&newTopic)
	} else {
		// Tópico já existe, então atualizamos as informações de retenção
		existingTopic.Name = publish.TopicName
		existingTopic.IdPublish = publish.Id
		existingTopic.Updated = time.Now()
		r.db.Save(&existingTopic)
	}

	c.JSON(http.StatusCreated, gin.H{"detail": "success.created.publish"})
}

// getAllPublishs obtém todas as publicações.
// @Summary Get all publishs
// @Description Get all publishs
// @Tags Publishs
// @Produce json
// @Success 200 {array}  models.PublishResponse
// @Router /publishs/all [get]
func (r *Routes) GetAllPublishs(c *gin.Context) {
	var publishsResponses []models.PublishResponse
	if err := r.db.Debug().Model(&models.Publish{}).
		Select("publishes.*, containers.* as \"Container\"").
		Joins("join containers on publishes.\"IdContainer\"=containers.\"Id\"").
		Scan(&publishsResponses).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting all topics"})
		r.logger.Error("Error: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, publishsResponses)
}

// getPublishByTopicName obtém uma publicação pelo ID.
// @Summary Get a publish by TopicName
// @Description Get a publish by TopicName
// @Tags Publishs
// @Produce json
// @Param TopicName query string true "Topic Name"
// @Success 200 {array}  models.PublishResponse
// @Router /publishs [get]
func (r *Routes) GetPublishByTopicName(c *gin.Context) {
	topicName := c.Query("TopicName")
	var publishsResponses []models.PublishResponse
	if err := r.db.Debug().Model(&models.Publish{}).
		Select("publishes.*, containers.* as \"Container\"").
		Joins("join containers on publishes.\"IdContainer\"=containers.\"Id\"").
		Where("publishes.\"TopicName\" = ?", topicName).
		Scan(&publishsResponses).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting all publish datas"})
		r.logger.Error("Error: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, publishsResponses)
}

// updatePublish atualiza uma publicação pelo ID.
// @Summary Update a publish by ID
// @Description Update a publish by ID
// @Tags Publishs
// @Accept json
// @Produce json
// @Param id path int true "Publish ID"
// @Param input body models.PublishRequest true "Updated publish object"
// @Success 200 {object} models.PublishRequest
// @Router /publishs/{id} [put]
func (r *Routes) UpdatePublish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var publish models.Publish
	if err := r.db.First(&publish, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&publish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Save(&publish)
	c.JSON(http.StatusOK, publish)
}
