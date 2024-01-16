package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
)

// createTopic cria um novo tópico.
// @Summary Create a new topic
// @Description Create a new topic
// @Tags Topic
// @Accept json
// @Produce json
// @Param input body TopicRequest true "Topic object that needs to be added"
// @Success 201 {object} Topic
// @Router /topics [post]
func (a *API) createTopic(c *gin.Context) {
	var topic Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Se o IdContainer existe, crie o tópico e associe-o ao contêiner
	var container Container
	if err := a.db.Where("\"Id\" = ?", topic.IdContainer).First(&container).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Container not found"})
		return
	}
	topic.IdContainer = uint64(container.Id)
	if err := a.db.Create(&topic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create topic"})
		return
	}

	c.JSON(http.StatusCreated, topic)
}

// getAllTopics obtém todos os tópicos.
// @Summary Get all topics
// @Description Get all topics
// @Tags Topic
// @Produce json
// @Success 200 {array} Topic
// @Router /topics [get]
func (a *API) getAllTopics(c *gin.Context) {
	var topicResponses []TopicResponse

	// Usando Preload para fazer um join e Select para escolher campos específicos
	if err := a.db.
		Model(&Topic{}).
		Select("topics.\"Id\", topics.\"Name\", topics.\"Payload\", topics.\"Qos\", topics.\"Created\", topics.\"Updated\", topics.\"Deleted\", containers.*").
		Joins("left join containers on topics.\"IdContainer\" = containers.\"Id\"").
		Where("topics.\"Deleted\" IS NULL").
		Scan(&topicResponses).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting all topics"})
		a.logger.Error("Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, topicResponses)
}

// getAllTopics obtém todos os tópicos.
// @Summary Get all topics
// @Description Get all topics
// @Tags Topic
// @Produce json
// @Param IdContainer path int true "Topic by IdContainer"
// @Success 200 {array} Topic
// @Router /topics/ByIdContainer/{IdContainer} [get]
func (a *API) getAllTopicsByIdContainer(c *gin.Context) {
	IdContainer, _ := strconv.Atoi(c.Param("IdContainer"))
	var topicResponses []TopicResponse

	// Usando Preload para fazer um join e Select para escolher campos específicos
	if err := a.db.
		Model(&Topic{}).
		Select("topics.\"Id\", topics.\"Name\", topics.\"Payload\", topics.\"Qos\", topics.\"Created\", topics.\"Updated\", topics.\"Deleted\", containers.*").
		Joins("left join containers on topics.\"IdContainer\" = containers.\"Id\"").
		Where(fmt.Sprintf("topics.\"IdContainer\"=%d", IdContainer)).
		Where("topics.\"Deleted\" IS NULL").
		Scan(&topicResponses).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting all topics"})
		a.logger.Error("Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, topicResponses)
}

// getTopicByID obtém um tópico pelo ID.
// @Summary Get a topic by ID
// @Description Get a topic by ID
// @Tags Topic
// @Produce json
// @Param id path int true "Topic ID"
// @Success 200 {object} Topic
// @Router /topics/{id} [get]
func (a *API) getTopicByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var topicResponse TopicResponse

	if err := a.db.
		Model(&Topic{}).
		Select("topics.\"Id\", topics.\"Name\", topics.\"Payload\", topics.\"Qos\", topics.\"Created\", topics.\"Updated\", topics.\"Deleted\", containers.*").
		Joins("left join containers on topics.\"IdContainer\" = containers.\"Id\"").
		Where("topics.\"Deleted\" IS NULL").
		First(&topicResponse, id).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting all topics"})
		a.logger.Error("Error: %s", err.Error())
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
// @Param input body TopicRequest true "Updated topic object"
// @Success 200 {object} TopicRequest
// @Router /topics/{id} [put]
func (a *API) updateTopic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var topic Topic
	if err := a.db.First(&topic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.db.Save(&topic)
	c.JSON(http.StatusOK, topic)
}
