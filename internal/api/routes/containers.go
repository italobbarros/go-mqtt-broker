package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/italobbarros/go-mqtt-broker/docs"
	models "github.com/italobbarros/go-mqtt-broker/internal/api/models"
)

// Example for Container model

// createContainer cria um novo container.
// @Summary Create a new container
// @Description Create a new container
// @Tags Containers
// @Accept json
// @Produce json
// @Param input body models.ContainerRequest true "Container object that needs to be added"
// @Success 201 {object} models.GenericResponse
// @Router /containers [post]
func (r *Routes) CreateContainer(c *gin.Context) {
	var container models.Container
	if err := c.ShouldBindJSON(&container); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.db.Create(&container)
	c.JSON(http.StatusCreated, container)
}

// getPublishByTopicName obtém uma publicação pelo ID.
// @Summary Get all containers info
// @Description Get all containers info
// @Tags Containers
// @Produce json
// @Success 200 {array}  models.ContainersInfoResponse
// @Router /containers/all [get]
func (r *Routes) GetAllContainers(c *gin.Context) {
	var containersInfoResponse []models.ContainersInfoResponse
	if err := r.db.
		Table("containers c").
		Select("c.\"Id\", COALESCE(COUNT(DISTINCT s.\"Id\"), 0) as \"CountSession\", COALESCE(COUNT(DISTINCT p.\"Id\"), 0) as \"CountPublishers\", COALESCE(COUNT(DISTINCT sub.\"Id\"), 0) as \"CountSubscribers\"").
		Joins("LEFT JOIN sessions s ON c.\"Id\" = s.\"IdContainer\"").
		Joins("LEFT JOIN publishes p ON s.\"Id\" = p.\"IdSession\"").
		Joins("LEFT JOIN subscriptions sub ON s.\"Id\" = sub.\"IdSession\"").
		Group("c.\"Id\"").
		Scan(&containersInfoResponse).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting container stats"})
		r.logger.Error("Error: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, containersInfoResponse)
}

// getContainerByID obtém um container pelo ID.
// @Summary Get a container by ID
// @Description Get a container by ID
// @Tags Containers
// @Produce json
// @Param id path int true "Container ID"
// @Success 200 {object} models.Container
// @Router /containers/{id} [get]
func (r *Routes) GetContainerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var container models.Container
	if err := r.db.First(&container, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, container)
}

// @Summary Delete a container by name
// @Description Delete a container by name
// @Tags Containers
// @Produce json
// @Param Name path string true "Container Name"
// @Success 204 {object} models.GenericResponse
// @Failure 400 {object} models.GenericResponse
// @Router /containers/{Name} [delete]
func (r *Routes) DeleteContainerByName(c *gin.Context) {
	// Extrair o nome do parâmetro da URL
	name := c.Param("Name")

	// Verificar se o container existe
	var container models.Container
	container.Name = name

	if err := r.db.Where("\"Name\" = ?", name).First(&container).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error getting container stats"})
		r.logger.Error("Error: %s", err.Error())
		return
	}

	// Excluir o container
	if err := r.db.Delete(&container).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Error deleting container"})
		r.logger.Error("Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"detail": "success.deleted"})
}
