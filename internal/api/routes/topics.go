package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	models "github.com/italobbarros/go-mqtt-broker/internal/api/models"
)

// getAllTopics obtém todos os tópicos.
// @Summary Get all topics
// @Description Get all topics
// @Tags Topic
// @Produce json
// @Success 200 {array} models.Topic
// @Router /topics/all [get]
func (r *Routes) GetAllTopics(c *gin.Context) {
	var topicResponses []models.TopicResponse

	// Usando Preload para fazer um join e Select para escolher campos específicos
	if err := r.db.
		Model(&models.Topic{}).
		Select("topics.\"Id\", topics.\"Name\", topics.\"Payload\", topics.\"Qos\", topics.\"Created\", topics.\"Updated\", topics.\"Deleted\", containers.*").
		Joins("left join publishes on topics.\"IdPublish\" = publishes.\"Id\"").
		Where("topics.\"Deleted\" IS NULL").
		Scan(&topicResponses).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting all topics"})
		r.logger.Error("Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, topicResponses)
}

// getAllTopics obtém todos os tópicos.
// @Summary Get Topic by Name
// @Description Get Topic by Name
// @Tags Topic
// @Produce json
// @Param Name query string true "Topic by Name"
// @Success 200 {object} models.Topic
// @Router /topics [get]
func (r *Routes) GetTopicsByName(c *gin.Context) {
	name := c.Query("Name")
	var topicResponse models.TopicResponse

	// Usando Preload para fazer um join e Select para escolher campos específicos
	if err := r.db.
		Model(&models.Topic{}).
		Select("topics.\"Id\", topics.\"Name\", topics.\"Created\", topics.\"Updated\", topics.\"Deleted\", publishes.*,containers.*").
		Joins("left join publishes on publishes.\"Id\"=topics.\"IdPublish\"").
		Joins("left join containers on containers.\"Id\"=publishes.\"IdContainer\"").
		Where("topics.\"Name\"=?", name).
		Where("topics.\"Deleted\" IS NULL").
		First(&topicResponse).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "error.get.topicsByName"})
		r.logger.Error("Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, topicResponse)
}

// getTopicByID obtém um tópico pelo ID.
// @Summary Get a topic by ID
// @Description Get a topic by ID
// @Tags Topic
// @Produce json
// @Param id path int true "Topic ID"
// @Success 200 {object} models.Topic
// @Router /topics/{id} [get]
func (r *Routes) GetTopicByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var topicResponse models.TopicResponse

	if err := r.db.
		Model(&models.Topic{}).
		Select("topics.\"Id\", topics.\"Name\", topics.\"Payload\", topics.\"Qos\", topics.\"Created\", topics.\"Updated\", topics.\"Deleted\", containers.*").
		Joins("left join publishes on topics.\"IdPublish\" = publishes.\"Id\"").
		Where("topics.\"Deleted\" IS NULL").
		First(&topicResponse, id).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting all topics"})
		r.logger.Error("Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, topicResponse)
}

// updateTopic atualiza um tópico pelo ID.
// @Summary Update a topic by ID
// @Description Update a topic by ID
// @Tags Topic
// @Accept json
// @Produce json
// @Param id path int true "Topic ID"
// @Param input body models.TopicRequest true "Updated topic object"
// @Success 200 {object} models.TopicRequest
// @Router /topics/{id} [put]
func (r *Routes) UpdateTopic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var topic models.Topic
	if err := r.db.First(&topic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Save(&topic)
	c.JSON(http.StatusOK, topic)
}
