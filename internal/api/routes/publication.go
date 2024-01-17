package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	models "github.com/italobbarros/go-mqtt-broker/internal/api/models"
)

// Endpoints for Publish

// createPublish cria uma nova publicação.
// @Summary Create a new publish
// @Description Create a new publish
// @Tags Publisher
// @Accept json
// @Produce json
// @Param input body models.PublishRequest true "Publish object that needs to be added"
// @Success 201 {object} models.GenericResponse
// @Router /publisher [post]
func (r *Routes) CreatePublish(c *gin.Context) {
	var publish models.Publish
	if err := c.ShouldBindJSON(&publish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publish.Timestamp = time.Now()
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

// getAllPublisher obtém todas as publicações.
// @Summary Get all publisher
// @Description Get all publisher
// @Tags Publisher
// @Produce json
// @Success 200 {array}  models.PublishResponse
// @Router /publisher/all [get]
func (r *Routes) GetAllPublisher(c *gin.Context) {
	var publisherResponses []models.PublishResponse
	if err := r.db.Debug().Model(&models.Publish{}).
		Select("publishes.*, sessions.*").
		Joins("join sessions on publishes.\"IdContainer\"=sessions.\"Id\"").
		Scan(&publisherResponses).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting all topics"})
		r.logger.Error("Error: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, publisherResponses)
}

// getPublishByTopicName obtém uma publicação pelo ID.
// @Summary Get a publish by TopicName
// @Description Get a publish by TopicName
// @Tags Publisher
// @Produce json
// @Param TopicName query string true "Topic Name"
// @Success 200 {array}  models.PublishResponse
// @Router /publisher/historic [get]
func (r *Routes) GetPublishByTopicName(c *gin.Context) {
	topicName := c.Query("TopicName")
	var publisherResponses []models.PublishResponse
	if err := r.db.Debug().Model(&models.Publish{}).
		Select("publishes.*, sessions.*").
		Joins("join sessions on publishes.\"IdSession\"=sessions.\"Id\"").
		Where("publishes.\"TopicName\" = ?", topicName).
		Scan(&publisherResponses).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting all publish datas"})
		r.logger.Error("Error: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, publisherResponses)
}
